package auth_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/cyrusn/ssgo/auth"
	jwt "github.com/dgrijalva/jwt-go"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("secret message"))
}
func in(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	type LoginForm struct {
		Username string
		Password string
	}
	errCode := http.StatusUnauthorized
	loginForm := new(LoginForm)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}
	if err := json.Unmarshal(body, loginForm); err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	username := loginForm.Username
	password := loginForm.Password

	for _, m := range testModels {
		if m.User.Username == username {
			claim := myClaims{
				Username: m.User.Username,
				Role:     m.User.Role,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expireToken,
				},
			}
			token, err := auth.Token(claim, m, m.User.Username, password)
			if err != nil {
				helper.PrintError(w, err, errCode)
				return
			}
			w.Write([]byte(token))
			return
		}
	}
	helper.PrintError(w, errors.New("User not found."), errCode)
}
