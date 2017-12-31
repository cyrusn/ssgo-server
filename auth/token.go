package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Authoriser is an interface that auth func
type Authoriser interface {
	Authorise(loginName, password string) error
}

func generateJWTToken(claim jwt.Claims, privateKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(privateKey)
}

// Token return jwt token if the auth success, where claim will be stored
// in jwt payload
func Token(claim jwt.Claims, auth Authoriser, loginName, password string) (string, error) {
	if err := auth.Authorise(loginName, password); err != nil {
		return "", err
	}

	return generateJWTToken(claim, key)
}
