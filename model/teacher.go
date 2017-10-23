package model

import (
	"golang.org/x/crypto/bcrypt"
)

// Teacher store information of teacher user
type Teacher struct {
	user.Info
}

// Insert insert teacher user in database
func (t *Teacher) Insert() error {

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

	if err != nil {
		return err
	}
	return nil
}

// Get get teacher information from database
func (t *Teacher) Get() error {
	statement := "SELECT * FROM teacher where username = ?"
	username := t.Username
	if err := db.QueryRow(statement, username).Scan(
		&teacher.Username,
		&teacher.Password,
		&teacher.Name,
		&teacher.Cname,
	); err != nil {
		return nil, err
	}

	return teacher, nil
}
