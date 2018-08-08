package route

import (
	"net/http"

	"github.com/cyrusn/ssgo/route/handler/auth"
	"github.com/cyrusn/ssgo/route/handler/student"
	"github.com/cyrusn/ssgo/route/handler/subject"
)

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
	return []Route{
		Route{
			// login request for student user
			Path:    "/auth/login",
			Methods: []string{"POST"},
			Scopes:  []string{},
			Auth:    false,
			Handler: auth.LoginHandler(env.Auth),
		},
		Route{
			Path:    "/auth/refresh/{jwtKeyName}",
			Methods: []string{"GET"},
			Scopes:  []string{},
			Auth:    true,
			Handler: auth.RefreshHandler(env.Auth),
		},
		Route{
			// get all students' status
			Path:    "/students",
			Methods: []string{"GET"},
			Scopes:  []string{"TEACHER", "ADMIN"},
			Auth:    true,
			Handler: student.ListHandler(env.Student),
		},
		Route{
			Path:    "/student/{userAlias}",
			Methods: []string{"GET"},
			Scopes:  []string{},
			Auth:    true,
			Handler: student.GetHandler(env.Student),
		},
		Route{
			// update student's priorities
			Path:    "/student/{userAlias}/priorities",
			Methods: []string{"PUT"},
			Scopes:  []string{"STUDENT"},
			Auth:    true,
			Handler: student.UpdatePrioritiesHandler(env.Student),
		},
		Route{
			// set student's isConfirmed value to true
			Path:    "/student/{userAlias}/isconfirmed/true",
			Methods: []string{"PUT"},
			Scopes:  []string{"STUDENT", "ADMIN"},
			Auth:    true,
			Handler: student.ConfirmedHandler(env.Student),
		},
		Route{
			// set student's isConfirmed value to false
			Path:    "/student/{userAlias}/isconfirmed/false",
			Methods: []string{"PUT"},
			Scopes:  []string{"ADMIN"},
			Auth:    true,
			Handler: student.UnconfirmedHandler(env.Student),
		},
		Route{
			// set student's rank value
			Path:    "/student/{userAlias}/rank/{rank}",
			Methods: []string{"PUT"},
			Scopes:  []string{"TEACHER", "ADMIN"},
			Auth:    true,
			Handler: student.UpdateRankHandler(env.Student),
		},
		Route{
			// list all subjects information
			Path:    "/subjects",
			Methods: []string{"GET"},
			Scopes:  []string{"TEACHER", "ADMIN"},
			Auth:    true,
			Handler: subject.ListHandler(env.Subject),
		},
		Route{
			// update subject's capacity
			Path:    "/subject/{subjectCode}/capacity/{capacity}",
			Methods: []string{"PUT"},
			Scopes:  []string{"TEACHER", "ADMIN"},
			Auth:    true,
			Handler: subject.UpdateCapacityHandler(env.Subject),
		},
	}
}
