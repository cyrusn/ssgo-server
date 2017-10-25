package model

type schema struct {
	name    string
	content string
}

var schemas = []schema{
	schema{"TEACHER", teacherTableSchema},
	schema{"STUDENT", studentTableSchema},
	schema{"SUBJECT", subjectTableSchema},
}

const teacherTableSchema = `
CREATE TABLE IF NOT EXISTS teacher (
  username TEXT PRIMARY KEY,
  password TEXT NOT NULL,
  name TEXT NOT NULL,
  cname TEXT
);`

const studentTableSchema = `
CREATE TABLE IF NOT EXISTS student (
  username TEXT PRIMARY KEY,
  password TEXT NOT NULL,
  name TEXT NOT NULL,
  cname TEXT,
	classcode TEXT NOT NULL,
	classno INTEGER NOT NULL,
	priority BLOB,
	is_confirmed INTEGER,
	rank INTEGER DEFAULT -1
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
func CreateTables() error {
	for _, s := range schemas {
		err := createTable(s.content)
		if err != nil {
			return err
		}
	}
	return nil
}

// createTable create table by given schema
func createTable(schema string) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
