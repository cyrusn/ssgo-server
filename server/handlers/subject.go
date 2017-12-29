package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cyrusn/ssgo/model"

	helper "github.com/cyrusn/goHTTPHelper"
)

// SubjectStore stores the interface for handler that query information about
// model.Subject
type SubjectStore interface {
	Get(subjectCode string) (*model.Subject, error)
	List() ([]*model.Subject, error)
	UpdateCapacity(subjectCode string, capacity int) error
}

// GetSubjectHandler get subject information by given subject code
func (env *Env) GetSubjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := env.Vars(r)
	subjectCode := vars["subjectCode"]
	subj, err := env.SubjectStore.Get(subjectCode)
	errCode := http.StatusBadRequest
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	helper.PrintJSON(w, subj, errCode)
}

// ListSubjectsHandler list all subjects' information
func (env *Env) ListSubjectsHandler(w http.ResponseWriter, r *http.Request) {
	subjects, err := env.SubjectStore.List()
	errCode := http.StatusBadRequest
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	helper.PrintJSON(w, subjects, errCode)
}

type capacityPostForm struct {
	Capacity int `json:"capacity"`
}

// UpdateSubjectCapacityHandler update subject capacity
func (env *Env) UpdateSubjectCapacityHandler(w http.ResponseWriter, r *http.Request) {
	vars := env.Vars(r)
	subjectCode := vars["subjectCode"]
	errCode := http.StatusBadRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.PrintError(w, err, errCode)
	}
	form := new(capacityPostForm)
	if err := json.Unmarshal(body, form); err != nil {
		helper.PrintError(w, err, errCode)
	}

	if err := env.SubjectStore.UpdateCapacity(subjectCode, form.Capacity); err != nil {
		helper.PrintError(w, err, errCode)
	}

	helper.PrintJSON(w, nil, errCode)
}
