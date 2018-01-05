package handlers

import (
	"net/http"
)

// Route stores information of a route in mux
type Route struct {
	Path    string
	Methods []string
	Scopes  []string
	Auth    bool
	Handler func(http.ResponseWriter, *http.Request)
}

// Routes return routes by given env
func Routes(env *Env) []Route {
	return []Route{
		Route{
			Path:    "/auth/students/login/",
			Methods: []string{"POST"},
			Scopes:  []string{},
			Auth:    false,
			Handler: env.StudentLoginHandler,
		},
		Route{
			Path:    "/auth/teachers/login/",
			Methods: []string{"POST"},
			Scopes:  []string{},
			Auth:    false,
			Handler: env.TeacherLoginHandler,
		},
		Route{
			Path:    "/auth/refresh/",
			Methods: []string{"GET"},
			Scopes:  []string{},
			Auth:    true,
			Handler: env.RefreshHandler,
		},
		Route{
			Path:    "/students/",
			Methods: []string{"GET"},
			Scopes:  []string{"TEACHER"},
			Auth:    true,
			Handler: env.ListStudentsHandler,
		},
		Route{
			Path:    "/students/{username}",
			Methods: []string{"GET"},
			Scopes:  []string{"STUDENT"},
			Auth:    true,
			Handler: env.GetStudentHandler,
		},
		Route{
			Path:    "/students/{username}/priority",
			Methods: []string{"PUT"},
			Scopes:  []string{"STUDENT", "TEACHER"},
			Auth:    true,
			Handler: env.UpdateStudentPriorityHandler,
		},
		Route{
			Path:    "/students/{username}/confirm",
			Methods: []string{"PUT"},
			Scopes:  []string{"STUDENT", "TEACHER"},
			Auth:    true,
			Handler: env.UpdateStudentIsConfirmedHandler,
		},
		Route{
			Path:    "/subjects/",
			Methods: []string{"GET"},
			Scopes:  []string{"TEACHER"},
			Auth:    true,
			Handler: env.ListSubjectsHandler,
		},
		Route{
			Path:    "/subjects/{subjectCode}",
			Methods: []string{"GET"},
			Scopes:  []string{"TEACHER"},
			Auth:    true,
			Handler: env.GetSubjectHandler,
		},
		Route{
			Path:    "/subjects/{subjectCode}/capacity",
			Methods: []string{"PUT"},
			Scopes:  []string{"TEACHER"},
			Auth:    true,
			Handler: env.UpdateSubjectCapacityHandler,
		},
	}
}
