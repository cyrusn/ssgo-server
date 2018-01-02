package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cyrusn/goHTTPHelper"

	jwt "github.com/dgrijalva/jwt-go"
)

// Scope is a middleware that parse jwt in header with value of roleKeyName as key
// (default value of roleKeyName is "Role", user can use SetRoleKeyName to set
// the vale of the roleKeyName).
// If value with roleKeyName in jwt payload is not in "scopes []string", handler will
// wthen print error message instead
func Scope(scopes []string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusUnauthorized
		role, err := parseRoleInContext(r.Context())
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		if ok := in(role, scopes); !ok {
			helper.PrintError(w, errors.New("User not allow"), errCode)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func parseRoleInContext(ctx context.Context) (string, error) {
	claim := ctx.Value(contextKeyName)
	if claim == nil {
		errMessage := fmt.Sprintf("%s not found in r.Context()", contextKeyName)
		return "", errors.New(errMessage)
	}
	m := claim.(jwt.MapClaims)
	result, ok := m[roleKeyName].(string)
	if !ok {
		errMessage := fmt.Sprintf("%s not found in %s(jwt.Claims)", roleKeyName, contextKeyName)
		return "", errors.New(errMessage)
	}

	return result, nil
}

func in(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}
