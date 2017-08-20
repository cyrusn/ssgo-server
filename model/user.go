package model

import "golang.org/x/crypto/bcrypt"

// User store basic information of teacher or student user
type User struct {
	Username  string
	Password  string
	Name      string
	Cname     string
	IsTeacher bool
}

// Validate validate the password of the user
func (u User) Validate(password string) error {
	bHash := []byte(u.Password)
	bPassword := []byte(password)
	return bcrypt.CompareHashAndPassword(bHash, bPassword)
}

// InsertUser insert a new user to user table
func (db *DB) InsertUser(u User) error {
	password := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
    INSERT INTO user (
      username, password, name, cname, is_teacher
    ) values (
      ?, ?, ?, ?, ?
    )`,
		u.Username, hash, u.Name, u.Cname, convertBool2Int(u.IsTeacher),
	)
	return err
}

// AllUsers queries all students
func (db *DB) AllUsers() ([]*User, error) {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []*User
	var isTeacher int

	for rows.Next() {
		u := new(User)
		err := rows.Scan(
			&u.Username, &u.Password, &u.Name, &u.Cname, &isTeacher,
		)

		if err != nil {
			return nil, err
		}
		u.IsTeacher = convertInt2Bool(isTeacher)
		users = append(users, u)
	}

	return users, nil
}

// GetUser query student by username
func (db *DB) GetUser(username string) (*User, error) {
	statement := "SELECT * FROM user where username = ?"

	row := db.QueryRow(statement, username)
	var isTeacher bool
	u := new(User)
	err := row.Scan(
		&u.Username, &u.Password, &u.Name, &u.Cname, &isTeacher,
	)
	if err != nil {
		return nil, err
	}

	u.IsTeacher = isTeacher

	return u, nil
}
