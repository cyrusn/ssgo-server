package handlers

import (
	"net/http"

	"github.com/cyrusn/ssgo/model"
)

// Env store all config and nessessary stores for retrieve data
type Env struct {
	Vars func(*http.Request) map[string]string
	StudentStore
}

// StudentStore stores the interface for handler that query information about model.Student
type StudentStore interface {
	Get(username string) (*model.Student, error)
	List() ([]*model.Student, error)
	UpdatePriority(username string, priority []int) error
	// UpdateIsConfirmed(string, bool)
}

// SubjectStore stores the interface for handler that query information about model.Subject
// type SubjectStore interface {
// 	Get(string) (*model.Subject, error)
// }
