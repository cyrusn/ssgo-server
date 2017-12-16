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
	"github.com/cyrusn/ssgo/server/handlers"
	"github.com/gorilla/mux"
)

type MockStudentStore struct {
	StudentList []*model.Student
}

var studentList = []*model.Student{
	&model.Student{
		User:        model.User{Username: "lpstudent1", Password: "password1", Name: "Alice Li", Cname: "李麗絲"},
		ClassCode:   "3A",
		ClassNo:     1,
		Priority:    []int{0, 1, 2, 3},
		IsConfirmed: false,
		Rank:        -1,
	},
	&model.Student{
		User:        model.User{Username: "lpstudent2", Password: "password2", Name: "Bob Li", Cname: "李鮑伯"},
		ClassCode:   "3B",
		ClassNo:     2,
		Priority:    []int{3, 2, 1, 0},
		IsConfirmed: false,
		Rank:        -1,
	},
	&model.Student{
		User:        model.User{Username: "lpstudent3", Password: "password3", Name: "Charlie Li", Cname: "李查利"},
		ClassCode:   "3C",
		ClassNo:     3,
		Priority:    []int{},
		IsConfirmed: true,
		Rank:        -1,
	},
}

var store = &MockStudentStore{studentList}
var env = &handlers.Env{
	StudentStore: store,
	Vars:         mux.Vars,
}

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

func (store *MockStudentStore) UpdatePriority(username string, priority []int) error {
	var counter = 0
	for _, s := range store.StudentList {
		if s.Username == username {
			s.Priority = priority
		} else {
			counter++
		}
	}

	if counter == len(store.StudentList)-1 {
		errorMessage := fmt.Sprintln("User not found: ", username)
		return errors.New(errorMessage)
	}

	return nil
}

func TestStudentHandlers(t *testing.T) {
	// Test Get for each student
	for i, student := range store.StudentList {
		name := fmt.Sprintf("GetStudentHandlers #%d", i+1)
		t.Run(name, testGetSudent(i, student))
	}
	// Test All
	t.Run("List all students", testListAllStudent)

	// Test UpdateStudentPriority
	t.Run("UpdateStudentPriority", testUpdateStudentPriority)
}

var testGetSudent = func(i int, student *model.Student) func(t *testing.T) {
	return func(t *testing.T) {
		w := httptest.NewRecorder()
		path := fmt.Sprintf("/students/lpstudent%d", i+1)
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
