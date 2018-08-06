package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/cyrusn/ssgo/model/student"
	"github.com/gorilla/mux"
)

// Store stores the interface for handler that query information about model.Student
type Store interface {
	Get(userAlias string) (*student.Student, error)
	List() ([]*student.Student, error)
	UpdateRank(userAlias string, rank int) error
	UpdatePriority(userAlias string, priority []int) error
	UpdateIsConfirmed(userAlias string, isConfirmed bool) error
}

// GetHandler get student information by given userAlias
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

// ListHandler get all students information
func ListHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := store.List()
		errCode := http.StatusBadRequest

		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		helper.PrintJSON(w, list)
	}
}

// UpdatePriorityHandler updated student's priority
func UpdatePriorityHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusBadRequest
		userAlias := mux.Vars(r)["userAlias"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		var form = new(struct {
			Priority []int `json:"priority"`
		})

		if err := json.Unmarshal(body, form); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		if err := store.UpdatePriority(userAlias, form.Priority); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		w.Write(nil)
	}
}

// ConfirmedHandler update IsConfirmed status of student
func ConfirmedHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return confirmHandlerBuilder(store, true)
}

// UnconfirmedHandler update IsConfirmed status of student
func UnconfirmedHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return confirmHandlerBuilder(store, false)
}

func confirmHandlerBuilder(store Store, b bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusBadRequest
		userAlias := mux.Vars(r)["userAlias"]

		if err := store.UpdateIsConfirmed(userAlias, b); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		w.Write(nil)
	}
}

// UpdateRankHandler update Rank of student
func UpdateRankHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusBadRequest
		userAlias := mux.Vars(r)["userAlias"]
		rank, err := strconv.Atoi(mux.Vars(r)["rank"])
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		if err := store.UpdateRank(userAlias, rank); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		w.Write(nil)
	}
}
