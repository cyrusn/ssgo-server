package model

import (
	"database/sql"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

// Student stores information for student user.
type Student struct {
	Info
	ClassCode   string
	ClassNo     int
	Priority    []int
	IsConfirmed bool
	Rank        int
}

type StudentList []*Student

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

// Insert add new student to database.
func (s *Student) Insert() error {
	priority, err := json.Marshal(s.Priority)
	if err != nil {
		return err
	}

	password := []byte(s.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO student (
			username,
			password,
			name,
			cname,
			classcode,
			classno,
			priority,
			is_confirmed,
			rank
		) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		s.Username,
		hashedPassword,
		s.Name,
		s.Cname,
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

// UpdatePriority will update student's priority.
func (s *Student) UpdatePriority(username string, p []int) error {
	priority, err := json.Marshal(p)
	if err != nil {
		return err
	}

	statement := "UPDATE student set priority = ? WHERE username = ?;"

	_, err = db.Exec(statement, priority, username)
	return err
}

// UpdateIsConfirmed will update student's isConfirmed.
func (s *Student) UpdateIsConfirmed(username string, b bool) error {
	statement := "UPDATE student set is_confirmed = ? WHERE username = ?;"

	_, err := db.Exec(statement, convertBool2Int(b), username)
	return err
}

// Get query student by username.
func (s *Student) Get(username string) (*Student, error) {
	statement := "SELECT * FROM student where username = ?"

	row := db.QueryRow(statement, username)
	err := s.scanStudent(row)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// All queries all students.
func (list StudentList) Get() (StudentList, error) {
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		s := new(Student)
		err := s.scanStudent(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}

	return list, nil
}

// scanStudent by *sql.Row or *sql.Rows.
func (s *Student) scanStudent(v interface{}) error {
	var priority []byte
	var isConfirmed int
	var err error

	var args = []interface{}{
		&s.Info.Username,
		&s.Info.Password,
		&s.Info.Name,
		&s.Info.Cname,
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
		return err
	}

	if err := json.Unmarshal(priority, &s.Priority); err != nil {
		return err
	}

	s.IsConfirmed = convertInt2Bool(isConfirmed)

	return nil
}
