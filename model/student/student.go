package student

import (
	"database/sql"
	"encoding/json"

	"github.com/cyrusn/ssgo/model/ssdb"
	"github.com/cyrusn/ssgo/model/user"

	"golang.org/x/crypto/bcrypt"
)

// Student stores information for student user.
type Student struct {
	user.Info
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

// Insert add new student to database.
func Insert(db *ssdb.DB, s Student) error {

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
func UpdatePriority(db *ssdb.DB, username string, p []int) error {
	priority, err := json.Marshal(p)
	if err != nil {
		return err
	}

	statement := "UPDATE student set priority = ? WHERE username = ?;"

	_, err = db.Exec(statement, priority, username)
	return err
}

// UpdateIsConfirmed will update student's isConfirmed.
func UpdateIsConfirmed(db *ssdb.DB, username string, b bool) error {
	statement := "UPDATE student set is_confirmed = ? WHERE username = ?;"

	_, err := db.Exec(statement, convertBool2Int(b), username)
	return err
}

// Get query student by username.
func Get(db *ssdb.DB, username string) (*Student, error) {
	statement := "SELECT * FROM student where username = ?"

	row := db.QueryRow(statement, username)
	return scanStudent(row)
}

// All queries all students.
func All(db *ssdb.DB) ([]*Student, error) {
	rows, err := db.Query("SELECT * FROM student")
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

// scanStudent by *sql.Row or *sql.Rows.
func scanStudent(v interface{}) (*Student, error) {
	s := new(Student)
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
		return nil, err
	}

	if err := json.Unmarshal(priority, &s.Priority); err != nil {
		return nil, err
	}

	s.IsConfirmed = convertInt2Bool(isConfirmed)

	return s, nil
}
