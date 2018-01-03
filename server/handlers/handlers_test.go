package handlers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/cyrusn/ssgo/model"
	"github.com/cyrusn/ssgo/server"
	"github.com/cyrusn/ssgo/server/handlers"

	auth "github.com/cyrusn/goJWTAuthHelper"
	"github.com/gorilla/mux"
)

var (
	routes       = server.Routes(env)
	r            = mux.NewRouter()
	studentStore = &MockStudentStore{studentList}
	sujectStore  = &MockSubjectStore{subjectList}
	teacherStore = &MockTeacherStore{teacherList}
	env          = &handlers.Env{
		StudentStore: studentStore,
		SubjectStore: sujectStore,
		TeacherStore: teacherStore,
		Vars:         mux.Vars,
	}
	mapToken        = make(map[string]string)
	contextClaimKey = "ssgo-claim"
	privateKey      = "hello-world"
	jwtKey          = "kid"
)

func init() {
	auth.SetPrivateKey(privateKey)
	auth.SetJWTKeyName(jwtKey)
	auth.SetContextKeyName(contextClaimKey)

	for _, route := range routes {
		handler := http.HandlerFunc(route.Handler)

		if len(route.Scopes) != 0 {
			handler = auth.Scope(route.Scopes, handler).(http.HandlerFunc)
		}

		if route.Auth {
			handler = auth.Validate(handler).(http.HandlerFunc)
		}

		r.Handle(route.Path, handler).Methods(route.Methods...)
	}
}

func TestMain(t *testing.T) {
	t.Run("Student Login", testStudentLogin)
	t.Run("Teacher login", testTeacherLogin)

	t.Run("Get Student", testGetStudent)
	t.Run("List all students", testListAllStudent)
	t.Run("UpdateStudentPriority", testUpdateStudentPriority)
	t.Run("UpdateStudentIsConfirmed", testUpdateStudentIsConfirmedHandler)

	t.Run("Get Subject", testGetSubjectHandler)
	t.Run("List Subjects", testListSubjectsHandler)
	t.Run("Update Subject Capacity", testUpdateSubjectCapacityHandler)
}

type MockStudentStore struct {
	StudentList []*model.Student
}

func (store *MockStudentStore) Authenticate(username, password string) error {
	s, err := store.Get(username)
	if err != nil {
		return err
	}
	if s.Password != password {
		return errors.New("Invalid password")
	}
	return nil
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

type MockSubjectStore struct {
	SubjectList []*model.Subject
}

func (store *MockSubjectStore) Get(subjectCode string) (*model.Subject, error) {
	for _, s := range store.SubjectList {
		if s.Code == subjectCode {
			return s, nil
		}
	}
	return nil, errors.New("Subject not found")
}

func (store *MockSubjectStore) List() ([]*model.Subject, error) {
	return store.SubjectList, nil
}

func (store *MockSubjectStore) UpdateCapacity(subjectCode string, capacity int) error {
	subj, err := store.Get(subjectCode)
	if err != nil {
		return err
	}
	subj.Capacity = capacity
	return nil
}

type MockTeacherStore struct {
	TeacherList []*model.Teacher
}

func (store *MockTeacherStore) Get(username string) (*model.Teacher, error) {
	for _, s := range store.TeacherList {
		if s.Username == username {
			return s, nil
		}
	}
	return nil, errors.New("Teacher not found")
}

func (store *MockTeacherStore) Authenticate(username, password string) error {
	s, err := store.Get(username)
	if err != nil {
		return err
	}
	if s.Password != password {
		return errors.New("Invalid password")
	}
	return nil
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

var teacherList = []*model.Teacher{
	&model.Teacher{
		Username: "lpteacher1", Password: "t-password1", Name: "Apple Li", Cname: "李蘋果",
	},
}

var subjectList = []*model.Subject{
	&model.Subject{
		Code:     "bio",
		Group:    1,
		Name:     "Biology",
		Cname:    "生物",
		Capacity: 0,
	},
	&model.Subject{
		Code:     "bafs",
		Group:    1,
		Name:     "Business, Accounting and Financial Studies",
		Cname:    "企業、會計與財務概論",
		Capacity: 0,
	},
	&model.Subject{
		Code:     "ict",
		Group:    2,
		Name:     "Information and Communication Technology",
		Cname:    "資訊及通訊科技",
		Capacity: 0,
	},
	&model.Subject{
		Code:     "econ",
		Group:    2,
		Name:     "Economics",
		Cname:    "經濟",
		Capacity: 0,
	},
}
