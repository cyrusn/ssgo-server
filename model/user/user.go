package user

import "golang.org/x/crypto/bcrypt"

// UserInfo store basic information of teacher or student user
type UserInfo struct {
	Username string
	Password string
	Name     string
	Cname    string
}

// Validate validate the password of the user
func (u UserInfo) Validate(password string) error {
	bHash := []byte(u.Password)
	bPassword := []byte(password)
	return bcrypt.CompareHashAndPassword(bHash, bPassword)
}
