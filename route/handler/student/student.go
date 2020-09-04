package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
	"ssgo-server/model/student"
	"github.com/gorilla/mux"
)

// Store stores the interface for handler that query information about model.Student
type Store interface {
	Get(userAlias string) (*student.Student, error)
	List() ([]*student.Student, error)
	UpdateRank(userAlias string, rank int) error
	UpdatePriorities(userAlias string, priorities []int) error
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

// UpdatePrioritiesHandler updated student's priorities
func UpdatePrioritiesHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusBadRequest
		userAlias := mux.Vars(r)["userAlias"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}
		var form = new(struct {
			Priorities []int `json:"priorities"`
		})

		if err := json.Unmarshal(body, form); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		if err := store.UpdatePriorities(userAlias, form.Priorities); err != nil {
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
		// userAlias := mux.Vars(r)["userAlias"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		var formDatas = []struct {
			UserAlias string `json:"userAlias"`
			Rank      int    `json:"rank"`
		}{}

		if err := json.Unmarshal(body, &formDatas); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		for _, d := range formDatas {
			if err := store.UpdateRank(d.UserAlias, d.Rank); err != nil {
				helper.PrintError(w, err, errCode)
				return
			}
		}
		w.Write(nil)
	}
}
