package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cyrusn/goHTTPHelper"
	"github.com/gorilla/mux"
)

type Store interface {
	Authenticate(userAlias, password string) (string, error)
	Refresh(token string) (string, error)
}

// LoginHandler get jwt token
func LoginHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userAlias, password, err := parseJSONPostForm(r)
		if err != nil {
			errCode := http.StatusBadRequest
			helper.PrintError(w, err, errCode)
			return
		}

		token, err := store.Authenticate(userAlias, password)

		if err != nil {
			errCode := http.StatusUnauthorized
			helper.PrintError(w, err, errCode)
			return
		}

		w.Write([]byte(token))
	}
}

// RefreshHandler refreshs the token
func RefreshHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusForbidden
		jwtKey := mux.Vars(r)["jwtKeyName"]
		token := r.Header.Get(jwtKey)

		newToken, err := store.Refresh(token)
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		w.Write([]byte(newToken))
	}
}

func parseJSONPostForm(r *http.Request) (string, string, error) {
	form := new(struct {
		UserAlias string
		Password  string
	})

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", "", err
	}

	if err = json.Unmarshal(body, form); err != nil {
		return "", "", err
	}

	return form.UserAlias, form.Password, nil
}
