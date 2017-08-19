package model

type schema struct {
	name    string
	content string
}

var schemas = []schema{
	schema{
		"USER", userTableSchema,
	},
	schema{
		"STUDENT", studentTableSchema,
	},
	schema{
		"SUBJECT", subjectTableSchema,
	},
}

const userTableSchema = `
CREATE TABLE IF NOT EXISTS user (
  username TEXT PRIMARY KEY,
  password TEXT NOT NULL,
  name TEXT NOT NULL,
  cname TEXT,
  is_teacher INTEGER NOT NULL
  );`

const studentTableSchema = `
CREATE TABLE IF NOT EXISTS students (
	username TEXT UNIQUE NOT NULL,
	classcode TEXT NOT NULL,
	classno INTEGER NOT NULL,
	priority BLOB,
	is_confirmed INTEGER,
	rank INTEGER,
	FOREIGN KEY(username) REFERENCES user(username)
);`

const subjectTableSchema = `
CREATE TABLE IF NOT EXISTS subject (
  code TEXT NOT NULL,
  gp INTEGER NOT NULL,
  name TEXT NOT NULL,
  cname TEXT NOT NULL,
  capacity INTEGER NOT NULL
);`

// CreateTables create all tables for ssgo system in database
func (db *DB) CreateTables() error {
	for _, s := range schemas {
		err := db.createTable(s.content)
		if err != nil {
			return err
		}
	}
	return nil
}

// createTable create table by given schema
func (db *DB) createTable(schema string) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
