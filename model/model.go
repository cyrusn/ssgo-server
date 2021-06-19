// Package model accept only mysql database
package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

type schema struct {
	name    string
	content string
}

const ERR_DATABASE_EXIST = "Database exsits"

const credentialTableSchema = `
CREATE TABLE IF NOT EXISTS Credential (
  userAlias varchar(64) PRIMARY KEY,
  password BLOB(128) NOT NULL,
  role varchar(64) NOT NULL
);`

const studentTableSchema = `
CREATE TABLE IF NOT EXISTS Student (
  userAlias varchar(64) PRIMARY KEY,
	priorities BLOB,
	isX3 BOOLEAN,
	isConfirmed BOOLEAN,
	ranking INTEGER DEFAULT 0,
	timestamp DATETIME NULL,
  FOREIGN KEY (userAlias) REFERENCES Credential(userAlias)
	);`

const subjectTableSchema = `
	CREATE TABLE IF NOT EXISTS Subject (
		code varchar(64) PRIMARY KEY,
		capacity INTEGER DEFAULT 0
	);`

const signatureTableSchema = `
	CREATE TABLE IF NOT EXISTS Signature (
		userAlias varchar(64) PRIMARY KEY,
  	isSigned BOOLEAN,
  	address MEDIUMTEXT,
		FOREIGN KEY (userAlias) REFERENCES Credential(userAlias)
	);`

var schemas = []schema{
	{"CREDENTIAL", credentialTableSchema},
	{"STUDENT", studentTableSchema},
	{"SUBJECT", subjectTableSchema},
	{"SIGNATURE", signatureTableSchema},
}

func ParseDSN(dsn string) (rootDSN, dbName string, err error) {
	config, err := mysql.ParseDSN(dsn)
	if err != nil {
		return "", "", err
	}
	dbName = config.DBName
	rootDSN = strings.Replace(dsn, dbName, "", 1)
	return rootDSN, dbName, nil
}

func IsDatabaseExist(db *sql.DB, dbName string) error {
	showDatabaseQuery := fmt.Sprintf("show databases like '%s'", dbName)
	if err := db.QueryRow(showDatabaseQuery).Scan(new(string)); err != nil {
		return err
	}
	return nil
}

func dropDatabase(db *sql.DB, dbName string) error {
	dropDatabaseQuery := fmt.Sprintf("drop database %s", dbName)
	_, err := db.Exec(dropDatabaseQuery)
	return err
}

func createDB(db *sql.DB, dbName string) error {
	createDBQuery := fmt.Sprintf("create database %s", dbName)
	_, err := db.Exec(createDBQuery)
	return err
}

// CreateDatabase create new database and write schema on it.
// To create long-lived db for server, use `sql.Open()` to open an database,
// and developer should check if the database is already existed.
func CreateDatabase(dsn string, isOverWrite bool) error {
	rootDSN, dbName, err := ParseDSN(dsn)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", rootDSN)
	defer db.Close()

	if err != nil {
		return err
	}

	if err := IsDatabaseExist(db, dbName); err != nil {
		return createDBAndTable(db, dbName)
	}

	if !isOverWrite {
		return errors.New(ERR_DATABASE_EXIST)
	}

	if err := dropDatabase(db, dbName); err != nil {
		return err
	}

	return createDBAndTable(db, dbName)
}

func createDBAndTable(db *sql.DB, dbName string) error {
	if err := createDB(db, dbName); err != nil {
		return err
	}

	useDBQuery := fmt.Sprintf("use %s", dbName)
	if _, err := db.Exec(useDBQuery); err != nil {
		return err
	}

	if err := createTable(db); err != nil {
		return err
	}
	return nil
}

// createTable create table for the schema of ScheduleSchema
func createTable(db *sql.DB) error {
	for _, schema := range schemas {
		if _, err := db.Exec(schema.content); err != nil {
			return err
		}
		fmt.Printf("Add %s schema.\n", schema.name)
	}
	return nil
}
