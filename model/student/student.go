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
	UserAlias   string `json:"userAlias"`
	Priorities  []int  `json:"priorities"`
	IsConfirmed bool   `json:"isConfirmed"`
	Rank        int    `json:"rank"`
}

// Insert add new student to database.
func (db *DB) Insert(s *Student) error {
	bPriorities, err := json.Marshal(s.Priorities)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO student (
			userAlias,
			priorities,
			isConfirmed,
			rank
		) values (?, ?, ?, ?)`,
		s.UserAlias,
		bPriorities,
		convertBool2Int(s.IsConfirmed),
		s.Rank,
	)

	if err != nil {
		return err
	}
	return nil
}

// UpdatePriorities will update student's priorities.
func (db *DB) UpdatePriorities(userAlias string, priorities []int) error {
	bPriorities, err := json.Marshal(priorities)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		"UPDATE student set priorities = ? WHERE (userAlias = ? and isConfirmed = false)",
		bPriorities,
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

// UpdateRank will update student's isConfirmed.
func (db *DB) UpdateRank(userAlias string, rank int) error {
	_, err := db.Exec(
		"UPDATE student set rank = ? WHERE userAlias = ?",
		rank,
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
