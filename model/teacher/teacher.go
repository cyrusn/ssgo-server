package teacher

import (
	"github.com/cyrusn/ssgo/model/ssdb"
	"github.com/cyrusn/ssgo/model/user"
	"golang.org/x/crypto/bcrypt"
)

// Teacher store information of teacher user
type Teacher user.Info

// Insert insert teacher user in database
func Insert(db *ssdb.DB, t Teacher) error {

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
func Get(db *ssdb.DB, username string) (*Teacher, error) {
	statement := "SELECT * FROM teacher where username = ?"

	var teacher = new(Teacher)
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
