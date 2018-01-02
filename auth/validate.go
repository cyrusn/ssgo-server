// Package auth is a auth server package which will handle the login request
// user will receive a JWT Token once sucessfull login. JWT token will
// contain basic information of the user which depend on application call.
package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/cyrusn/goHTTPHelper"
	jwt "github.com/dgrijalva/jwt-go"
)

// Validate is a middleware which will check if jwt in request header is valid
func Validate(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get(jwtKeyName)
		if jwtToken == "" {
			errCode := http.StatusForbidden
			err := errors.New("Token not found")
			helper.PrintError(w, err, errCode)
			return
		}

		token, err := jwt.Parse(jwtToken, keyFunc)
		errCode := http.StatusUnauthorized
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), contextKeyName, claims)
			req := r.WithContext(ctx)
			handler.ServeHTTP(w, req)
			return
		}
		helper.PrintError(w, errors.New("Invalid Token"), errCode)
	})
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return privateKey, nil
}
