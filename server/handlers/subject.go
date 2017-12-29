package handlers

import (
	"net/http"

	"github.com/cyrusn/ssgo/model"

	helper "github.com/cyrusn/goHTTPHelper"
)

// SubjectStore stores the interface for handler that query information about
// model.Subject
type SubjectStore interface {
	Get(subjectCode string) (*model.Subject, error)
	List() ([]*model.Subject, error)
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
