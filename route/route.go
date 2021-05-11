package route

import (
	"net/http"

	"ssgo-server/route/handler/auth"
	"ssgo-server/route/handler/student"
	"ssgo-server/route/handler/subject"
)

// Env contain stores for providing values to handlers
type Env struct {
	Auth    auth.Store
	Student student.Store
	Subject subject.Store
}

// Route stores information of a route in mux
type Route struct {
	Path    string
	Methods []string
	Scopes  []string
	Auth    bool
	Handler func(http.ResponseWriter, *http.Request)
}

// Routes return routes by given env
func (env *Env) Routes() []Route {
	return []Route{{
		// login request for student user
		Path:    "/auth/login",
		Methods: []string{"POST"},
		Scopes:  []string{},
		Auth:    false,
		Handler: auth.LoginHandler(env.Auth),
	}, {
		Path:    "/auth/refresh/{jwtKeyName}",
		Methods: []string{"GET"},
		Scopes:  []string{},
		Auth:    true,
		Handler: auth.RefreshHandler(env.Auth),
	}, {
		// get all students' status
		Path:    "/students",
		Methods: []string{"GET"},
		Scopes:  []string{"TEACHER", "ADMIN"},
		Auth:    true,
		Handler: student.ListHandler(env.Student),
	}, {
		Path:    "/student/{userAlias}",
		Methods: []string{"GET"},
		Scopes:  []string{"STUDENT"},
		Auth:    true,
		Handler: student.GetHandler(env.Student),
	}, {
		// update student's priorities
		Path:    "/student/{userAlias}/priorities",
		Methods: []string{"PUT"},
		Scopes:  []string{"STUDENT"},
		Auth:    true,
		Handler: student.UpdatePrioritiesHandler(env.Student),
	}, {
		// update student's olePriorities
		Path:    "/student/{userAlias}/ole_priorities",
		Methods: []string{"PUT"},
		Scopes:  []string{"STUDENT"},
		Auth:    true,
		Handler: student.UpdateOlePrioritiesHandler(env.Student),
	}, {
		// set student's isConfirmed value to true
		Path:    "/student/{userAlias}/isconfirmed/true",
		Methods: []string{"PUT"},
		Scopes:  []string{"STUDENT", "ADMIN"},
		Auth:    true,
		Handler: student.ConfirmedHandler(env.Student),
	}, {
		// set student's isConfirmed value to false
		Path:    "/student/{userAlias}/isconfirmed/false",
		Methods: []string{"PUT"},
		Scopes:  []string{"ADMIN"},
		Auth:    true,
		Handler: student.UnconfirmedHandler(env.Student),
	}, {
		// set student's rank value
		Path:    "/students/rank",
		Methods: []string{"PUT"},
		Scopes:  []string{"ADMIN"},
		Auth:    true,
		Handler: student.UpdateRankHandler(env.Student),
	}, {
		// list all subjects information
		Path:    "/subjects",
		Methods: []string{"GET"},
		Scopes:  []string{"ADMIN"},
		Auth:    true,
		Handler: subject.ListHandler(env.Subject),
	}, {
		// update subject's capacity
		Path:    "/subject/{subjectCode}/capacity/{capacity}",
		Methods: []string{"PUT"},
		Scopes:  []string{"ADMIN"},
		Auth:    true,
		Handler: subject.UpdateCapacityHandler(env.Subject),
	}}
}
