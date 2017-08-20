package server

import "net/http"

type Route struct {
	Path    string
	Methods []string
	Scope   []string
	Handler func(http.ResponseWriter, *http.Request)
}

var routes = []Routes{
	Route{
		"/auth/login",
		[]string{"GET"},
		[]string{"STUDENT", "TEACHER", "ADMIN"},
		loginHandler,
	},
	Route{
		"/auth/refresh",
		[]string{"GET"},
		[]string{"STUDENT", "TEACHER", "ADMIN"},
		refreshHandler,
	},
	Route{
		"/user/details",
		[]string{"GET"},
		[]string{"STUDENT", "TEACHER", "ADMIN"},
		userDetailHandler,
	},
	Route{
		"/student/priority",
		[]string{"GET", "POST"},
		[]string{"STUDENT"},
		priorityHandler,
	},
	Route{
		"/student/confirm",
		[]string{"POST"},
		[]string{"STUDENT"},
		studentConfirmHandler
	},
	Route{
		"/teacher/confirm/toggle/{studentID}",
		[]string{"POST"},
		[]string{"TEACHER", "ADMIN"},
		teacherToggleConfirmHandler,
	},
	Route{
		"/subject/capacity",
		[]string{"GET", "POST"},
		[]string{"ADMIN"},
		capacityHandler,
	},
	Route{
		"/subject/allocation",
		[]string{"POST"},
		[]string{"ADMIN"},
		allocationHandler,
	},
}
