package subject

import (
	"net/http"
	"strconv"

	helper "github.com/cyrusn/goHTTPHelper"
	"ssgo-server/model/subject"
	"github.com/gorilla/mux"
)

// Store stores the interface for handler that query information about model.Subject
type Store interface {
	List() ([]*subject.Subject, error)
	UpdateCapacity(subjectCode string, capacity int) error
}

// ListHandler list all subjects' information
func ListHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusBadRequest

		subjects, err := store.List()
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		helper.PrintJSON(w, subjects)
	}
}

// UpdateCapacityHandler update subject capacity
func UpdateCapacityHandler(store Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errCode := http.StatusBadRequest

		subjectCode := mux.Vars(r)["subjectCode"]

		capacity, err := strconv.Atoi(mux.Vars(r)["capacity"])
		if err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		if err := store.UpdateCapacity(subjectCode, capacity); err != nil {
			helper.PrintError(w, err, errCode)
			return
		}

		w.Write(nil)
	}
}
