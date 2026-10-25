package main

import (
	"encoding/json"
	"fmt"
	"lder/pkg/db"
	"lder/pkg/tmdb"
	"lder/pkg/tvdb"
	"lder/router"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	Port     uint      `json:"port"`
	TmdbKey  string    `json:"tmdb_key"`
	TvdbKey  string    `json:"tvdb_key"`
	Database db.Config `json:"database"`
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error getting current working directory")
	}
	configPath := filepath.Join(cwd, "/config.json")
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file %s\n", configPath)
	}
	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config file %s\n", err)
	}
	if config.Port == 0 {
		log.Fatalln("Invalid port provided")
	}

	tmdb := tmdb.New(config.TmdbKey)
	tvdb := tvdb.New(config.TvdbKey)
	if !tmdb.ValidateKey() {
		log.Fatalln("Invalid TMDB key")
	}
	if !tvdb.ValidateKey() {
		log.Fatalln("Invalid TVDB key")
	}

	database, err := db.NewDbAfterMigrations(
		config.Database,
		filepath.Join(cwd, "migrations"),
	)
	if err != nil {
		log.Fatalln("Error initialising DB ", err)
	}
	defer database.Close()

	httpRouter := router.New()
	handler := httpRouter.SetupRoutes()
	fmt.Println("Starting http server")
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handler)
	if err != nil {
		log.Fatalf("Error starting http server %s\n", err)
	}
}
