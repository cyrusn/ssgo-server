package model

import (
	"database/sql"
	"encoding/json"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile + log.LstdFlags)
}

// Student stores information for users including studnt, teacher and admin
type Student struct {
	Username    string
	Password    string
	Name        string
	Cname       string
	Priority    []int
	IsConfirmed bool
}

const studentTableSchema = `
CREATE TABLE IF NOT EXISTS students (
	username TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	name TEXT NOT NULL,
	cname TEXT NOT NULL,
	priority BLOB,
	isConfirmed INTEGER
);`

func convertBool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func convertInt2Bool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

// CreateStudentTable create student table
func (db *DB) CreateStudentTable() error {
	return db.createTable(studentTableSchema)
}

// AddStudent add new student to database
func (db *DB) AddStudent(s Student) error {

	priority, err := json.Marshal(s.Priority)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO students (
			username,
			password,
			name,
			cname,
			priority,
			isConfirmed
		) values (?, ?, ?, ?, ?, ?)`,
		s.Username,
		s.Password,
		s.Name,
		s.Cname,
		priority,
		convertBool2Int(s.IsConfirmed),
	)

	return err
}

// UpdatePriorityInStudentsTable will update student's priority
func (db *DB) UpdatePriorityInStudentsTable(username string, p []int) error {
	priority, err := json.Marshal(p)
	if err != nil {
		return err
	}

	statement := "UPDATE students set priority = ? WHERE username = ?;"

	_, err = db.Exec(statement, priority, username)
	return err
}

// UpdateIsConfirmedInStudentsTable will update student's isConfirmed
func (db *DB) UpdateIsConfirmedInStudentsTable(username string, b bool) error {
	statement := "UPDATE students set isConfirmed = ? WHERE username = ?;"

	_, err := db.Exec(statement, convertBool2Int(b), username)
	return err
}

// GetStudent query student by username
func (db *DB) GetStudent(username string) (*Student, error) {
	statement := "SELECT * FROM students where username = ?"

	row := db.QueryRow(statement, username)
	return scanStudent(row)
}

// AllStudents queries all students
func (db *DB) AllStudents() ([]*Student, error) {
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []*Student

	for rows.Next() {
		s, err := scanStudent(rows)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

// scanStudent by *sql.Row or *sql.Rows
func scanStudent(v interface{}) (*Student, error) {
	s := new(Student)
	var priority []byte
	var isConfirmed int
	var err error

	var args = []interface{}{
		&s.Username,
		&s.Password,
		&s.Name,
		&s.Cname,
		&priority,
		&isConfirmed,
	}

	switch t := v.(type) {
	case *sql.Row:
		err = t.Scan(args...)
	case *sql.Rows:
		err = t.Scan(args...)
	}

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(priority, &s.Priority); err != nil {
		return nil, err
	}

	s.IsConfirmed = convertInt2Bool(isConfirmed)
	return s, nil
}
