package model

import (
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

// CreateStudentTable create student table
func (db *DB) CreateStudentTable() error {
	const studentTableSchema = `
	CREATE TABLE IF NOT EXISTS students (
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		name TEXT NOT NULL,
		cname TEXT NOT NULL,
		priority BLOB,
		isConfirmed INTEGER
	);`

	_, err := db.Exec(studentTableSchema)
	if err != nil {
		return err
	}
	log.Println("STUDENTS table created")
	return nil
}

// AllStudents queries all students from database
func (db *DB) AllStudents() ([]*Student, error) {
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []*Student

	for rows.Next() {
		s := new(Student)
		var priority []byte
		var isConfirmed int

		err := rows.Scan(
			&s.Username,
			&s.Password,
			&s.Name,
			&s.Cname,
			&priority,
			&isConfirmed,
		)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(priority, &s.Priority); err != nil {
			return nil, err
		}

		s.IsConfirmed = convertInt2Bool(isConfirmed)

		students = append(students, s)
	}

	return students, nil
}
