package handlers

import (
	"net/http"
)

// GetStudentHandler get student information by given username
func (env *Env) GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
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
	list, err := env.StudentListStore.Get()
	errCode := http.StatusBadRequest

	if err != nil {
		errPrint(w, err, errCode)
		return
	}
	jsonPrint(w, list, errCode)
}

// UpdateStudentPriorityHandler updated student's priority
// func (env *Env) UpdateStudentPriorityHandler(w http.ResponseWriter, r *http.Request) {
// 	errCode := http.StatusBadRequest
// 	username := r.FormValue("username")
// 	priorityString := r.PostFormValue("id")
// 	fmt.Println(priorityString)
// 	priorityByte := []byte(priorityString)
//
// 	var priority []int
// 	if err := json.Unmarshal(priorityByte, &priority); err != nil {
// 		errPrint(w, err, errCode)
// 	}
// 	err := env.StudentStore.UpdatePriority(username, priority)
// 	if err != nil {
// 		errPrint(w, err, errCode)
// 		return
// 	}
// 	jsonPrint(w, nil, errCode)
// }
