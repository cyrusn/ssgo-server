package handlers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"

	"testing"

	"github.com/cyrusn/ssgo/helper"
	"github.com/cyrusn/ssgo/model"
	"github.com/cyrusn/ssgo/server/handlers"
)

type MockStudentStore struct {
	StudentList []*model.Student
}

var studentList = []*model.Student{
	&model.Student{model.User{"lpstudent1", "password1", "Alice Li", "李麗絲"}, "3A", 1, []int{0, 1, 2, 3}, false, -1},
	&model.Student{model.User{"lpstudent2", "password2", "Bob Li", "李鮑伯"}, "3A", 2, []int{3, 2, 1, 0}, false, -1},
	&model.Student{model.User{"lpstudent3", "password3", "Charlie Li", "李查利"}, "3A", 3, []int{}, true, -1},
}
var store = &MockStudentStore{studentList}
var env = &handlers.Env{StudentStore: store}

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
	// Test Get
	for i, student := range store.StudentList {
		w := httptest.NewRecorder()
		name := fmt.Sprintf("GetStudentHandlers #%d", i+1)

		t.Run(name, func(t *testing.T) {
			url := fmt.Sprintf("/api/get/student?username=lpstudent%d", i+1)
			req := httptest.NewRequest("Get", url, nil)

			env.GetStudentHandler(w, req)
			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			got := new(model.Student)
			if err := json.Unmarshal(body, got); err != nil {
				t.Fatal(err)
			}

			want := student
			helper.DiffTest(got, want, t)
		})
	}

	// Test All
	t.Run("AllStudents", func(t *testing.T) {
		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/get/student/all")
		req := httptest.NewRequest("Get", url, nil)
		env.AllStudentsHandler(w, req)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		var got []*model.Student

		want := studentList
		if err := json.Unmarshal(body, &got); err != nil {
			t.Fatal(err)
		}

		helper.DiffTest(got, want, t)
	})

	// Test UpdateStudentPriority
	// t.Run("UpdateStudentPriority", func(t *testing.T) {
	// 	// var result = []int{2, 1, 3, 0}
	// 	for i, sts := range store.StudentList {
	// 		w := httptest.NewRecorder()
	//
	// 		data := url.Values{"key": {"Value"}, "id": {"123"}}
	//
	// 		urlstring := fmt.Sprintf("/api/get/student/priority/update?username=lpstudent%d", i+1)
	// 		req := httptest.NewRequest("POST", urlstring, strings.NewReader(data.Encode()))
	//
	// 		env.UpdateStudentPriorityHandler(w, req)
	//
	// 		fmt.Println(sts.Priority)
	// 	}
	// })
}
