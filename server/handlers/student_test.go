package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/cyrusn/ssgo/model"
	"github.com/gorilla/mux"
)

func TestStudentHandlers(t *testing.T) {
	// Test Get for each student
	t.Run("Get Student", testGetStudent)

	// Test All
	t.Run("List all students", testListAllStudent)

	// Test UpdateStudentPriority
	t.Run("UpdateStudentPriority", testUpdateStudentPriority)

	// Test UpdateStudentIsConfirme
	t.Run("UpdateStudentIsConfirmed", testUpdateStudentIsConfirm)
}

var testGetStudent = func(t *testing.T) {
	for _, student := range studentList {
		w := httptest.NewRecorder()
		path := fmt.Sprintf("/students/%s", student.User.Username)
		req := httptest.NewRequest("GET", path, nil)

		r := mux.NewRouter()
		r.HandleFunc("/students/{username}", env.GetStudentHandler)
		r.ServeHTTP(w, req)

		resp := w.Result()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := new(model.Student)
		if err := json.Unmarshal(body, got); err != nil {
			t.Fatal(err)
		}
		want := student
		diffTest(got, want, t)
	}
}

var testListAllStudent = func(t *testing.T) {
	w := httptest.NewRecorder()

	url := fmt.Sprintf("/students/list")
	req := httptest.NewRequest("Get", url, nil)

	r := mux.NewRouter()
	r.HandleFunc("/students/list", env.AllStudentsHandler)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var got []*model.Student

	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}

	diffTest(got, studentList, t)
}

var testUpdateStudentPriority = func(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	for _, s := range studentList {
		url := fmt.Sprintf("/students/%s/priority", s.Username)
		r.HandleFunc("/students/{username}/priority", env.UpdateStudentPriorityHandler)

		form := strings.NewReader(`{"priority":[0, 1, 2, 3]}`)
		req := httptest.NewRequest("PUT", url, form)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
	}

	for _, s := range studentList {
		diffTest(s.Priority, []int{0, 1, 2, 3}, t)
	}
}

var testUpdateStudentIsConfirm = func(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	for _, s := range studentList {
		url := fmt.Sprintf("/student/%s/isConfirm", s.Username)
		r.HandleFunc("/student/{username}/isConfirm", env.UpdateStudentIsConfirmHandler)
		form := strings.NewReader(`{"isConfirmed":true}`)
		req := httptest.NewRequest("PUT", url, form)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
	}
	for _, s := range studentList {
		diffTest(s.IsConfirmed, true, t)
	}
}
