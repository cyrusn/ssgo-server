package handlers

import "github.com/cyrusn/ssgo/model"

// TeacherStore stores the interface for handler that query information about model.Teacher
type TeacherStore interface {
	Authenticate(username, password string) error
	Get(username string) (*model.Teacher, error)
}
