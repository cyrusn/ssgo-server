package student

import (
	"database/sql"
	"encoding/json"
)

type DB struct {
	*sql.DB
}

// Student stores information for student user.
type Student struct {
	UserAlias   string
	ClassCode   string
	ClassNo     int
	Priority    []int
	IsConfirmed bool
	Rank        int
}

// Insert add new student to database.
func (db *DB) Insert(s *Student) error {
	bPriority, err := json.Marshal(s.Priority)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO student (
			userAlias,
			classcode,
			classno,
			priority,
			isConfirmed,
			rank
		) values (?, ?, ?, ?, ?, ?)`,
		s.UserAlias,
		s.ClassCode,
		s.ClassNo,
		bPriority,
		convertBool2Int(s.IsConfirmed),
		s.Rank,
	)

	if err != nil {
		return err
	}
	return nil
}

// UpdatePriority will update student's priority.
func (db *DB) UpdatePriority(userAlias string, priority []int) error {
	bPriority, err := json.Marshal(priority)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		"UPDATE student set priority = ? WHERE userAlias = ?",
		bPriority,
		userAlias,
	)
	return err
}

// UpdateIsConfirmed will update student's isConfirmed.
func (db *DB) UpdateIsConfirmed(userAlias string, isConfirmed bool) error {
	_, err := db.Exec(
		"UPDATE student set isConfirmed = ? WHERE userAlias = ?",
		convertBool2Int(isConfirmed),
		userAlias,
	)
	return err
}

// Get query student by userAlias.
func (db *DB) Get(userAlias string) (*Student, error) {
	s := new(Student)
	row := db.QueryRow("SELECT * FROM Student where userAlias = ?", userAlias)
	err := s.scanStudent(row)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// List get all students.
func (db *DB) List() ([]*Student, error) {
	var list []*Student
	rows, err := db.Query("SELECT * FROM Student")
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
