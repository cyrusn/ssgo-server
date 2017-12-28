package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/cyrusn/ssgo/model"
	"github.com/gorilla/mux"
)

type route struct {
	url     string
	methods []string
	handler func(http.ResponseWriter, *http.Request)
}

var routes = []route{
	route{
		url:     "/students/",
		methods: []string{"GET"},
		handler: env.ListStudentsHandler,
	},
	route{
		url:     "/students/{username}",
		methods: []string{"GET"},
		handler: env.GetStudentHandler,
	},
	route{
		url:     "/students/{username}/priority",
		methods: []string{"PUT"},
		handler: env.UpdateStudentPriorityHandler,
	},
	route{
		url:     "/students/{username}/confirm",
		methods: []string{"PUT"},
		handler: env.UpdateStudentIsConfirmedHandler,
	},
}

var r = mux.NewRouter()

func init() {
	for _, route := range routes {
		r.HandleFunc(route.url, route.handler).Methods(route.methods...)
	}
}

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
		want := student
		diffTest(got, want, t)
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

	diffTest(got, studentList, t)
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
		diffTest(s.Priority, []int{0, 1, 2, 3}, t)
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
		diffTest(s.IsConfirmed, true, t)
	}
}
