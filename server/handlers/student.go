package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/cyrusn/ssgo/model"
)

// StudentStore stores the interface for handler that query information about model.Student
type StudentStore interface {
	Get(username string) (*model.Student, error)
	List() ([]*model.Student, error)
	UpdatePriority(username string, priority []int) error
	UpdateIsConfirmed(username string, isConfirmed bool) error
}

type priorityPostForm struct {
	Priority []int
}

type isConfirmedPostForm struct {
	IsConfirmed bool
}

// GetStudentHandler get student information by given username
func (env *Env) GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := env.Vars(r)
	username := vars["username"]
	s, err := env.StudentStore.Get(username)

	errCode := http.StatusBadRequest
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}
	helper.PrintJSON(w, s, errCode)
	return
}

// ListStudentsHandler get all students information
func (env *Env) ListStudentsHandler(w http.ResponseWriter, r *http.Request) {
	list, err := env.List()
	errCode := http.StatusBadRequest

	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}
	helper.PrintJSON(w, list, errCode)
}

// UpdateStudentPriorityHandler updated student's priority
func (env *Env) UpdateStudentPriorityHandler(w http.ResponseWriter, r *http.Request) {
	errCode := http.StatusBadRequest
	vars := env.Vars(r)
	username := vars["username"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	var form = new(priorityPostForm)
	if err := json.Unmarshal(body, form); err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	if err := env.StudentStore.UpdatePriority(username, form.Priority); err != nil {
		helper.PrintError(w, err, errCode)
		return
	}
	helper.PrintJSON(w, nil, errCode)
}

// UpdateStudentIsConfirmedHandler update IsConfirmed status of student
func (env *Env) UpdateStudentIsConfirmedHandler(w http.ResponseWriter, r *http.Request) {
	errCode := http.StatusBadRequest
	vars := env.Vars(r)
	username := vars["username"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	var form = new(isConfirmedPostForm)

	if err := json.Unmarshal(body, form); err != nil {
		helper.PrintError(w, err, errCode)
		return
	}
	if err := env.StudentStore.UpdateIsConfirmed(username, form.IsConfirmed); err != nil {
		helper.PrintError(w, err, errCode)
	}
}
