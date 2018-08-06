// Package model accept only sqlite3 database
package model

import (
	"database/sql"
	"fmt"
	"os"
)

type schema struct {
	name    string
	content string
}

const credentialTableSchema = `
CREATE TABLE IF NOT EXISTS Credential (
  userAlias TEXT PRIMARY KEY,
  password BLOB NOT NULL,
  role TEXT NOT NULL,
  FOREIGN KEY (userAlias) REFERENCES User(userAlias)
);`

const studentTableSchema = `
CREATE TABLE IF NOT EXISTS Student (
  userAlias TEXT PRIMARY KEY,
	priority BLOB,
	isConfirmed INTEGER,
	rank INTEGER DEFAULT -1,
  FOREIGN KEY (userAlias) REFERENCES User(userAlias)
);`

const subjectTableSchema = `
CREATE TABLE IF NOT EXISTS Subject (
  code TEXT PRIMARY KEY,
  capacity INTEGER DEFAULT -1
);`

var schemas = []schema{
	schema{"CREDENTIAL", credentialTableSchema},
	schema{"STUDENT", studentTableSchema},
	schema{"SUBJECT", subjectTableSchema},
}

// CreateDBFile create new database and write schema on it.
// To create long-lived db for server, use `sql.Open()` to open an database,
// and developer should check if the database is already existed.
func CreateDBFile(dbPath string, isOverWrite bool) error {
	exist, err := IsFileExist(dbPath)
	if err != nil {
		return err
	}

	if exist && !isOverWrite {
		return os.ErrExist
	}

	if _, err := os.Create(dbPath); err != nil {
		return err
	}

	if err := openDBAndCreateTable(dbPath); err != nil {
		return err
	}

	return nil
}

// openDBAndCreateTable open existed database and create table
func openDBAndCreateTable(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := createTable(db, dbPath); err != nil {
		return err
	}

	return nil
}

// createTable create table for the schema of ScheduleSchema
func createTable(db *sql.DB, dbPath string) error {
	for _, schema := range schemas {
		if _, err := db.Exec(schema.content); err != nil {
			return err
		}
		fmt.Printf("Add %s schema.\n", schema.name)
	}
	return nil
}

// IsFileExist check if file exist
func IsFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// file exist
		return true, nil

	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
