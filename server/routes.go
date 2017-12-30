package server

import (
	"net/http"

	"github.com/cyrusn/ssgo/server/handlers"
)

// Route stores information of a route in mux
type Route struct {
	Path    string
	Methods []string
	// Scope   []string
	Handler func(http.ResponseWriter, *http.Request)
}

// Routes return routes by given env
func Routes(env *handlers.Env) []Route {
	return []Route{
		Route{
			Path:    "/students/",
			Methods: []string{"GET"},
			Handler: env.ListStudentsHandler,
		},
		Route{
			Path:    "/students/{username}",
			Methods: []string{"GET"},
			Handler: env.GetStudentHandler,
		},
		Route{
			Path:    "/students/{username}/priority",
			Methods: []string{"PUT"},
			Handler: env.UpdateStudentPriorityHandler,
		},
		Route{
			Path:    "/students/{username}/confirm",
			Methods: []string{"PUT"},
			Handler: env.UpdateStudentIsConfirmedHandler,
		},
		Route{
			Path:    "/subjects/",
			Methods: []string{"GET"},
			Handler: env.ListSubjectsHandler,
		},
		Route{
			Path:    "/subjects/{subjectCode}",
			Methods: []string{"GET"},
			Handler: env.GetSubjectHandler,
		},
		Route{
			Path:    "/subjects/{subjectCode}/capacity",
			Methods: []string{"PUT"},
			Handler: env.UpdateSubjectCapacityHandler,
		},
	}
}
