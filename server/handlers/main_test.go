package handlers_test

import (
	"errors"

	"github.com/cyrusn/ssgo/model"
	"github.com/cyrusn/ssgo/server/handlers"
	"github.com/gorilla/mux"
)

var env = &handlers.Env{
	StudentStore: store,
	Vars:         mux.Vars,
}

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
