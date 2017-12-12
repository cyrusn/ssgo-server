package model

import (
	"golang.org/x/crypto/bcrypt"
)

// Teacher store information of teacher user
type Teacher User

// Insert insert teacher user in database
func (db TeacherDB) Insert(t *Teacher) error {

	password := []byte(t.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO teacher (
      username,
      password,
      name,
      cname
    ) values (?, ?, ?, ?)`,
		t.Username,
		hashedPassword,
		t.Name,
		t.Cname,
	)

	return err
}

// Get get teacher information from database
func (db TeacherDB) Get(username string) (*Teacher, error) {
	statement := "SELECT * FROM teacher where username = ?"
	t := new(Teacher)
	if err := db.QueryRow(statement, username).Scan(
		&t.Username,
		&t.Password,
		&t.Name,
		&t.Cname,
	); err != nil {
		return nil, err
	}

	return t, nil
}