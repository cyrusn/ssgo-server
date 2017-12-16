package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cyrusn/ssgo/model"
)

// StudentStore stores the interface for handler that query information about model.Student
type StudentStore interface {
	Get(username string) (*model.Student, error)
	List() ([]*model.Student, error)
	UpdatePriority(username string, priority []int) error
	// UpdateIsConfirmed(username string, isConfirm bool) error
}

// GetStudentHandler get student information by given username
func (env *Env) GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := env.Vars(r)
	username := vars["username"]
	s, err := env.StudentStore.Get(username)

	errCode := http.StatusBadRequest
	if err != nil {
		errPrint(w, err, errCode)
		return
	}
	jsonPrint(w, s, errCode)
	return
}

// AllStudentsHandler get all students information
func (env *Env) AllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	list, err := env.List()
	errCode := http.StatusBadRequest

	if err != nil {
		errPrint(w, err, errCode)
		return
	}
	jsonPrint(w, list, errCode)
}

type priorityPostForm struct {
	Priority []int
}

// UpdateStudentPriorityHandler updated student's priority
func (env *Env) UpdateStudentPriorityHandler(w http.ResponseWriter, r *http.Request) {
	errCode := http.StatusBadRequest
	vars := env.Vars(r)
	username := vars["username"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errPrint(w, err, errCode)
		return
	}

	var form = new(priorityPostForm)
	if err := json.Unmarshal(body, form); err != nil {
		errPrint(w, err, errCode)
		return
	}

	if err := env.StudentStore.UpdatePriority(username, form.Priority); err != nil {
		errPrint(w, err, errCode)
		return
	}
	jsonPrint(w, nil, errCode)
}
