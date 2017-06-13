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

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// createTable create table by given schema
func (db *DB) createTable(schema string) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
