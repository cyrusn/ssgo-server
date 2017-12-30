package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/model"
)

func TestSubjectHandler(t *testing.T) {
	t.Run("Get Subject", testGetSubjectHandler)
	t.Run("List Subjects", testListSubjectsHandler)
	t.Run("Update Subject Capacity", testUpdateSubjectCapacityHandler)
}

var testGetSubjectHandler = func(t *testing.T) {
	for _, subj := range subjectList {
		url := fmt.Sprintf("/subjects/%s", subj.Code)
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", url, nil)
		r.ServeHTTP(w, req)

		resp := w.Result()

		body, err := ioutil.ReadAll(resp.Body)
		assert.OK(t, err)

		got := new(model.Subject)
		assert.OK(t, json.Unmarshal(body, got))
		assert.Equal(got, subj, t)
	}
}

var testListSubjectsHandler = func(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/subjects/", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()

	body, err := ioutil.ReadAll(resp.Body)
	assert.OK(t, err)
	var got []*model.Subject
	assert.OK(t, json.Unmarshal(body, &got))

	for i, subj := range got {
		assert.Equal(subj, subjectList[i], t)
	}
}

var testUpdateSubjectCapacityHandler = func(t *testing.T) {
	for _, subj := range subjectList {
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/subjects/%s/capacity", subj.Code)
		form := strings.NewReader(`{"capacity":20}`)
		req := httptest.NewRequest("PUT", url, form)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
	}

	for _, s := range subjectList {
		assert.Equal(s.Capacity, 20, t)
	}
}
