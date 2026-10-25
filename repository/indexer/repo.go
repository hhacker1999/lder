package indexer_repo

import (
	"database/sql"
	"fmt"
)

type Indexer struct {
	Id               int    `json:"-"`
	Name             string `json:"name"`
	Host             string `json:"host"`
	ApiKey           string `json:"api_key"`
	ApiPath          string `json:"api_path"`
	AdditionalParams string `json:"additional_params"`
	Priority         int    `json:"priority"`
	Enabled          bool   `json:"enabled"`
}

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) InsertIndexer(indexer Indexer) error {
	query := `insert into newznab_indexers name, host, api_key, api_path, additional_params, priority, enabled
	values ($1, $2, $3, $4, $5, %6, %7)
	`
	_, err := r.db.Exec(
		query,
		indexer.Name,
		indexer.Host,
		indexer.ApiKey,
		indexer.ApiPath,
		indexer.AdditionalParams,
		indexer.Priority,
		indexer.Enabled,
	)
	if err != nil {
		fmt.Println("Error inserting into newznab indexers", err)
	}

	return err
}

func (r *Repo) RemoveIndexer(id int) error {
	query := `delete from newznab_indexers where id = $1
	`
	_, err := r.db.Exec(
		query,
		id,
	)
	if err != nil {
		fmt.Println("Error removing from newznab indexers", err)
	}

	return err
}

func (r *Repo) GetAllIndexers() ([]Indexer, error) {
	var res []Indexer
	query := `select id, name, host, api_key, api_path, additional_params, priority, enabled
	from newznab_indexers
	`
	rows, err := r.db.Query(
		query,
	)
	if err != nil {
		fmt.Println("Error inserting into newznab indexers", err)
	}
	defer rows.Close()
	for rows.Next() {
		indexer := Indexer{}
		apiKey := sql.NullString{}
		apiPath := sql.NullString{}
		apiParam := sql.NullString{}
		err = rows.Scan(
			&indexer.Id,
			&indexer.Name,
			&indexer.Host,
			&apiKey,
			&apiPath,
			&apiParam,
			&indexer.Priority,
			&indexer.Enabled,
		)
		if err != nil {
			fmt.Println("Error scanning indexer", err)
			return res, err
		}
		if apiKey.Valid {
			indexer.ApiKey = apiKey.String
		}
		if apiParam.Valid {
			indexer.AdditionalParams = apiParam.String
		}
		if apiPath.Valid {
			indexer.ApiPath = apiPath.String
		}
		res = append(res, indexer)
	}

	return res, err
}
