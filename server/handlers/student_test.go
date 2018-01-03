package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/model"
)

var testGetStudent = func(t *testing.T) {
	for _, s := range studentList {
		path := fmt.Sprintf("/students/%s", s.User.Username)
		req := httptest.NewRequest("GET", path, nil)

		addJWT2Header(s.Username, req)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		body := parseBody(w, t)

		got := new(model.Student)
		if err := json.Unmarshal(body, got); err != nil {
			t.Fatal(err)
		}

		assert.Equal(got, s, t)
	}
}

var testListAllStudent = func(t *testing.T) {
	req := httptest.NewRequest("GET", "/students/", nil)

	addJWT2Header(teacherList[0].Username, req)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	body := parseBody(w, t)

	var got []*model.Student
	assert.OK(t, json.Unmarshal(body, &got))
	assert.Equal(got, studentList, t)
}

var testUpdateStudentPriority = func(t *testing.T) {
	for _, s := range studentList {
		url := fmt.Sprintf("/students/%s/priority", s.Username)

		form := strings.NewReader(`{"priority":[0, 1, 2, 3]}`)
		req := httptest.NewRequest("PUT", url, form)
		addJWT2Header(s.Username, req)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}

	for _, s := range studentList {
		assert.Equal(s.Priority, []int{0, 1, 2, 3}, t)
	}
}

var testUpdateStudentIsConfirmedHandler = func(t *testing.T) {
	for _, s := range studentList {
		url := fmt.Sprintf("/students/%s/confirm", s.Username)
		form := strings.NewReader(`{"isConfirmed":true}`)
		req := httptest.NewRequest("PUT", url, form)
		req.Header.Set("Content-Type", "application/json")

		addJWT2Header(s.Username, req)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}
	for _, s := range studentList {
		assert.Equal(s.IsConfirmed, true, t)
	}
}

func addJWT2Header(username string, req *http.Request) {
	token := mapToken[username]
	req.Header.Set(jwtKey, token)
}

func parseBody(w *httptest.ResponseRecorder, t *testing.T) []byte {
	resp := w.Result()
	if resp.StatusCode >= 400 {
		assert.OK(t, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.OK(t, err)
	return body
}
