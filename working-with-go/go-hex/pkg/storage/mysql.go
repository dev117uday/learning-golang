package storage

import "database/sql"

type Storage struct {
	db *sql.DB
}

func SetupStorage() (*Storage, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Password02117 dbname=learn")
	if err != nil {
		return nil, sql.ErrConnDone
	}
	return &Storage{db: db}, nil
}
