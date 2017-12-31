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

// Required is a middleware for http.Handler that need right to access
func Required(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get(roleKeyName)
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