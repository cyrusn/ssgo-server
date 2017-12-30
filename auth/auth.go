// Package auth is a auth server package which will handle the login request
// user will receive a JWT Token once sucessfull login. JWT token will
// contain basic information of the user which depend on application call.
package auth

import (
	"context"
	"errors"
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
	jwt "github.com/dgrijalva/jwt-go"
)

// Claim store information of JWT token
type Claim struct {
	Payload interface{}
	jwt.StandardClaims
}

type privateKey []byte
type contextKey string

func (c contextKey) String() string {
	return "auth context key " + string(c)
}

var (
	key = []byte("secret")
	// ContextKeyClaim is the key in context for retreive the information of claim
	ContextKeyClaim = contextKey("claims")
)

// SetPrivateKey set the privateKey for authentication
func SetPrivateKey(newKey []byte) {
	key = newKey
}

// Required is a middleware for http.Handler that need right to access
func Required(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("jwt-token")
		if jwtToken == "" {
			errCode := http.StatusForbidden
			helper.PrintError(w, errors.New("Token not found"), errCode)
			return
		}

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		errCode := http.StatusUnauthorized
		if err != nil {
			helper.PrintError(w, err, errCode)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), ContextKeyClaim, claims)
			req := r.WithContext(ctx)
			handler.ServeHTTP(w, req)
		} else {
			helper.PrintError(w, errors.New("Invalid Token"), errCode)
		}
	})
}

// Authoriser is an interface the auth func
type Authoriser interface {
	Authorise(username string, password string) error
}

func (claim *Claim) generateJWTToken(privateKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(privateKey)
}

// GetToken send jwt token if the Basic Auth success
func (claim *Claim) GetToken(auth Authoriser, username, password string) (string, error) {
	err := auth.Authorise(username, password)
	if err != nil {
		return "", err
	}

	return claim.generateJWTToken(key)
}
