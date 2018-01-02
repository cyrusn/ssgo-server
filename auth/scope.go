package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/cyrusn/goHTTPHelper"

	jwt "github.com/dgrijalva/jwt-go"
)

// Scope is a middleware that parse jwt in header with key "Role", if value of
// "Role" in jwt payload is not in "scope", []string, handler then will
// print error message instead
func Scope(scope []string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusUnauthorized
		role, err := parseRoleInContext(r.Context())
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		if ok := in(role, scope); !ok {
			helper.PrintError(w, errors.New("User not allow"), errCode)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func parseRoleInContext(ctx context.Context) (string, error) {
	claim := ctx.Value(contextKeyName)
	if claim == nil {
		return "", errors.New("Claim not found in context")
	}
	m := claim.(jwt.MapClaims)
	result, ok := m[roleKeyName].(string)
	if !ok {
		return "", errors.New("Role not found in context")
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
