package signature

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ssgo-server/model/signature"
	"strconv"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/gorilla/mux"
)

type Store interface {
	Get(userAlias string) (*signature.Signature, error)
	UpdateIsSigned(userAlias string, isSigned bool) error
	UpdateAddress(userAlias string, signature string) error
}

// GetHandler get signature information by given userAlias
func GetHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userAlias := vars["userAlias"]
		s, err := store.Get(userAlias)

		errCode := http.StatusBadRequest
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		helper.PrintJSON(w, s)
		return
	}
}

func UpdateAddressHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusBadRequest
		userAlias := mux.Vars(r)["userAlias"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		var form = new(struct {
			Signature string `json:"signature"`
		})

		if err := json.Unmarshal(body, form); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		if err := store.UpdateAddress(userAlias, form.Signature); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		w.Write(nil)
	}
}

// UpdateIsSignedHandler update IsSigned status of Signature
func UpdateIsSignedHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusBadRequest
		userAlias := mux.Vars(r)["userAlias"]
		strBool := mux.Vars(r)["bool"]
		b, err := strconv.ParseBool(strBool)
		if err != nil {
			errCode := http.StatusBadRequest
			helper.PrintError(w, err, errCode)
			fmt.Println(userAlias, b, "damn")
			return
		}
		if err := store.UpdateIsSigned(userAlias, b); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		w.Write(nil)
	}
}
