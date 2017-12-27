package handlers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"strings"

	"testing"

	"github.com/cyrusn/ssgo/model"
	"github.com/gorilla/mux"
)

type MockStudentStore struct {
	StudentList []*model.Student
}

var store = &MockStudentStore{studentList}

func (store *MockStudentStore) Get(username string) (*model.Student, error) {
	for _, s := range store.StudentList {
		if s.Username == username {
			return s, nil
		}
	}
	return nil, errors.New("User not found")
}

func (store *MockStudentStore) List() ([]*model.Student, error) {
	return studentList, nil
}

func (store *MockStudentStore) UpdateIsConfirmed(username string, isConfirmed bool) error {
	student, err := store.Get(username)
	if err != nil {
		return err
	}

	student.IsConfirmed = isConfirmed
	return nil
}

func (store *MockStudentStore) UpdatePriority(username string, priority []int) error {
	student, err := store.Get(username)
	if err != nil {
		return err
	}
	student.Priority = priority
	return nil
}

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

// expectError is a testing tool, it used to test for error handling
func expectError(name string, t *testing.T, f func()) {
	defer func(t *testing.T) {
		err := recover()

		if err == nil {
			t.Fatalf("Error Test: [%s] did not return error", name)
		}
	}(t)
	f()
}

// diffTest is simply test if there are differences of 2 structs
func diffTest(got, want interface{}, t *testing.T) {
	if !reflect.DeepEqual(want, got) {

		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			got,
			want,
		)
	}
}
