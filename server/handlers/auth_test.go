package handlers_test

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cyrusn/goTestHelper"
)

var testStudentLogin = func(t *testing.T) {
	for _, student := range studentList {
		w := httptest.NewRecorder()
		formString := fmt.Sprintf(`{"Username":"%s", "Password":"%s"}`, student.Username, student.Password)
		postForm := strings.NewReader(formString)
		req := httptest.NewRequest("POST", "/auth/students/login/", postForm)
		r.ServeHTTP(w, req)

		resp := w.Result()

		body, err := ioutil.ReadAll(resp.Body)
		assert.OK(t, err)

		mapToken[student.Username] = string(body)
	}
}

var testTeacherLogin = func(t *testing.T) {
	for _, teacher := range teacherList {
		w := httptest.NewRecorder()
		formString := fmt.Sprintf(`{"Username":"%s", "Password":"%s"}`, teacher.Username, teacher.Password)
		postForm := strings.NewReader(formString)
		req := httptest.NewRequest("POST", "/auth/teachers/login/", postForm)
		r.ServeHTTP(w, req)

		resp := w.Result()

		body, err := ioutil.ReadAll(resp.Body)
		assert.OK(t, err)

		mapToken[teacher.Username] = string(body)
	}
}
