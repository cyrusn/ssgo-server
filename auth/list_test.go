package auth_test

import "net/http"

type User struct {
	Username string
	Role     string
	Password string
}

type Credential struct {
	Username string
	Password string
}

type TestModel struct {
	Name string
	*User
	*Credential
}

var testModels = []*TestModel{
	&TestModel{
		Name: "Success Case",
		User: &User{
			Username: "lpteacher1",
			Password: "abc123",
			Role:     "TEACHER",
		},
		Credential: &Credential{
			Username: "lpteacher1",
			Password: "abc123",
		},
	},
	&TestModel{
		Name: "Incorrect Login",
		User: &User{
			Username: "lpstudent1",
			Password: "def456",
			Role:     "STUDENT",
		},
		Credential: &Credential{
			Username: "lpstudent1",
			Password: "def123",
		},
	},
	&TestModel{
		Name: "Forbidden Role",
		User: &User{
			Username: "lpstudent2",
			Password: "ghi789",
			Role:     "STUDENT",
		},
		Credential: &Credential{
			Username: "lpstudent2",
			Password: "ghi789",
		},
	},
}

type route struct {
	path    string
	auth    bool
	roles   []string
	method  string
	handler func(http.ResponseWriter, *http.Request)
}

var routes = []route{
	route{
		path:    "/login/",
		auth:    false,
		roles:   []string{},
		method:  "GET",
		handler: loginHandler,
	},
	route{
		path:    "/auth/",
		auth:    true,
		roles:   authorizedRoles,
		method:  "GET",
		handler: simpleHandler,
	},
	route{
		path:    "/basic/",
		auth:    false,
		roles:   []string{},
		method:  "GET",
		handler: simpleHandler,
	},
}
