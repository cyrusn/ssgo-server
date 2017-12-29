package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"testing"

	helper "github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/model"
)

func TestStudentHandlers(t *testing.T) {
	// Test Get for each student
	t.Run("Get Student", testGetStudent)

	// Test All
	t.Run("List all students", testListAllStudent)

	// Test UpdateStudentPriority
	t.Run("UpdateStudentPriority", testUpdateStudentPriority)

	// Test UpdateStudentIsConfirme
	t.Run("UpdateStudentIsConfirmed", testUpdateStudentIsConfirmedHandler)
}

var testGetStudent = func(t *testing.T) {
	for _, student := range studentList {
		path := fmt.Sprintf("/students/%s", student.User.Username)
		req := httptest.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
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

		helper.Diff(got, student, t)
	}
}

var testListAllStudent = func(t *testing.T) {
	req := httptest.NewRequest("GET", "/students/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var got []*model.Student
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}

	helper.Diff(got, studentList, t)
}

var testUpdateStudentPriority = func(t *testing.T) {
	for _, s := range studentList {
		url := fmt.Sprintf("/students/%s/priority", s.Username)

		form := strings.NewReader(`{"priority":[0, 1, 2, 3]}`)
		req := httptest.NewRequest("PUT", url, form)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}

	for _, s := range studentList {
		helper.Diff(s.Priority, []int{0, 1, 2, 3}, t)
	}
}

var testUpdateStudentIsConfirmedHandler = func(t *testing.T) {
	for _, s := range studentList {
		url := fmt.Sprintf("/students/%s/confirm", s.Username)
		form := strings.NewReader(`{"isConfirmed":true}`)
		req := httptest.NewRequest("PUT", url, form)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}
	for _, s := range studentList {
		helper.Diff(s.IsConfirmed, true, t)
	}
}
