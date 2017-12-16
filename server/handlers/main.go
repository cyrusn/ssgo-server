package handlers

import (
	"net/http"
)

// Env store all config and nessessary stores for retrieve data
type Env struct {
	Vars func(*http.Request) map[string]string
	Port string
	StudentStore
}

// SubjectStore stores the interface for handler that query information about model.Subject
// type SubjectStore interface {
// 	Get(string) (*model.Subject, error)
// }
