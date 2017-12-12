package handlers_test

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/cyrusn/ssgo/helper"
// 	"github.com/cyrusn/ssgo/model"
// 	"github.com/cyrusn/ssgo/server/handlers"
// )

// var subjectList = []model.Subject{
// 	model.Subject{"bio", 1, "Biology", "生物", 0},
// 	model.Subject{"bafs", 1, "Business, Accounting and Financial Studies", "企業、會計與財務概論", 0},
// 	model.Subject{"ict", 2, "Information and Communication Technology", "資訊及通訊科技", 0},
// 	model.Subject{"econ", 2, "Economics", "經濟", 0},
// }

// func (subjectDB) GetSubject(code string) (*model.Subject, error) {
// 	for _, sub := range subjectList {
// 		if sub.Code == code {
// 			return &sub, nil
// 		}
// 	}
// 	return nil, errors.New("Subject not found")
// }

// func TestSubject(t *testing.T) {
// 	db := new(MockDB)
// 	env := &handlers.Env{Datastore: db}

// 	for i, s := range subjectList {
// 		rec := httptest.NewRecorder()
// 		name := fmt.Sprintf("GetSubject #%d", i+1)
// 		t.Run(name, func(t *testing.T) {
// 			url := fmt.Sprintf("/api/get/subject?code=%s", s.Code)
// 			req, _ := http.NewRequest("Get", url, nil)

// 			http.HandlerFunc(env.GetSubjectHandler).ServeHTTP(rec, req)
// 			want := &s

// 			got := new(*model.Subject)
// 			body := rec.Body
// 			if err := json.Unmarshal(body.Bytes(), got); err != nil {
// 				t.Fatal(err)
// 			}

// 			helper.DiffTest(got, &want, t)
// 		})
// 	}
// }
