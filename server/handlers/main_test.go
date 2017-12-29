package handlers_test

import (
	"errors"

	"github.com/cyrusn/ssgo/model"
	"github.com/cyrusn/ssgo/server"
	"github.com/cyrusn/ssgo/server/handlers"
	"github.com/gorilla/mux"
)

var (
	routes       = server.Routes(env)
	r            = mux.NewRouter()
	studentStore = &MockStudentStore{studentList}
	sujectStore  = &MockSubjectStore{subjectList}
	env          = &handlers.Env{
		StudentStore: studentStore,
		SubjectStore: sujectStore,
		Vars:         mux.Vars,
	}
)

func init() {
	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Methods...)
	}
}

type MockStudentStore struct {
	StudentList []*model.Student
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
