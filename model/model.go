package model

import (
	"database/sql"
)

var db *sql.DB

// InitDB startup DB for model package
func InitDB(dbPath string) {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	// Ping also establish a connection if necessary
	if err = db.Ping(); err != nil {
		panic(err)
	}
}
