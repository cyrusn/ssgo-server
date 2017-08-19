package model

import (
	"database/sql"
	"encoding/json"
)

// Student stores information for users including studnt, teacher and admin
type Student struct {
	Username    string
	ClassCode   string
	ClassNo     int
	Priority    []int
	IsConfirmed bool
	Rank        int
}

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

// InsertStudent add new student to database
func (db *DB) InsertStudent(s Student) error {

	priority, err := json.Marshal(s.Priority)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO students (
			username,
			classcode,
			classno,
			priority,
			is_confirmed,
			rank
		) values (?, ?, ?, ?, ?, ?)`,
		s.Username,
		s.ClassCode,
		s.ClassNo,
		priority,
		convertBool2Int(s.IsConfirmed),
		s.Rank,
	)

	if err != nil {
		return err
	}
	return nil
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
	statement := "UPDATE students set is_confirmed = ? WHERE username = ?;"

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
		&s.ClassCode,
		&s.ClassNo,
		&priority,
		&isConfirmed,
		&s.Rank,
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
