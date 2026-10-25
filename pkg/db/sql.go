package db

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func NewDbAfterMigrations(cfg Config, migrationsDirectory string) (*sql.DB, error) {
	if cfg.User == "" {
		return nil, fmt.Errorf("database user is required")
	}

	dsn := postgresDSN(cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	database.SetMaxOpenConns(25)
	database.SetMaxIdleConns(5)
	database.SetConnMaxLifetime(30 * time.Minute)

	if err := database.Ping(); err != nil {
		database.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	if err := runMigrations(dsn, migrationsDirectory); err != nil {
		database.Close()
		return nil, err
	}

	return database, nil
}

func runMigrations(dsn string, migrationsDirectory string) error {
	absDir, err := filepath.Abs(migrationsDirectory)
	if err != nil {
		return fmt.Errorf("resolve migrations path: %w", err)
	}

	sourceURL := migrationSourceURL(absDir)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("open postgres for migrations: %w", err)
	}
	defer database.Close()

	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	if err != nil {
		return fmt.Errorf("create migrator: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
}

func migrationSourceURL(absDir string) string {
	slashPath := filepath.ToSlash(absDir)
	if strings.HasPrefix(slashPath, "/") {
		return "file://" + slashPath
	}
	return "file://" + slashPath
}

func postgresDSN(user, password, host string, port uint, dbName string) string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   dbName,
	}
	q := u.Query()
	q.Set("sslmode", "disable")
	u.RawQuery = q.Encode()
	return u.String()
}
