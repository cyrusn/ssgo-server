package handlers

import (
	"net/http"
)

// Env store all config and nessessary stores for retrieve data
type Env struct {
	// Vars stores the key value pair information from *http.Request
	// which is used to parse the variables on api
	Vars func(*http.Request) map[string]string
	Port string
	StudentStore
	SubjectStore
}
