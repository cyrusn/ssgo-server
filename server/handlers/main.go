package handlers

import "github.com/cyrusn/ssgo/model"

// StudentStore stores the interface for handler that query information about model.Student
type StudentStore interface {
	Get(username string) (*model.Student, error)
	UpdatePriority(username string, priority []int) error
	// UpdateIsConfirmed(string, bool)
}

type StudentListStore interface {
	Get() (model.StudentList, error)
}

// SubjectStore stores the interface for handler that query information about model.Subject
// type SubjectStore interface {
// 	Get(string) (*model.Subject, error)
// }

// Env ...
type Env struct {
	StudentStore
	StudentListStore
	// SubjectStore
	// other config...
}
