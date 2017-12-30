// Package auth is a auth server package which will handle the login request
// user will receive a JWT Token once sucessfull login. JWT token will
// contain basic information of the user which depend on application call.
package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Username string
	Role     string
	jwt.StandardClaims
}

// Validate validate JWT Token
func Validate(jwtToken string) error {
	return nil
}

// Validate create JWT token for user
func GenerateJWTToken(u *User) (string, error) {
	key := "hello world"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Sign() {}

func ParseJWT() {}
