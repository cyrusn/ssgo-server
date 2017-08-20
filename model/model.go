package model

import (
	"database/sql"
)

// DB is custom database for model
type DB struct {
	*sql.DB
}

// InitDB startup DB for model package
func InitDB(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Ping also establish a connection if necessary
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
