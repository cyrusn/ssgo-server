package model

import (
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	log.SetFlags(log.Llongfile)
}

// InitDB startup DB for model package
func InitDB(dbPath string) {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Panic(err)
	}

	// Ping also establish a connection if necessary
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
