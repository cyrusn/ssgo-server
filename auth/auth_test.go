package auth_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/auth"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

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

var testRoutes = []route{
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

const (
	roleKeyName      = "kid"
	contextClaimName = "myClaim"
	privateKey       = "hello world"
)

var (
	r               = mux.NewRouter()
	expireToken     = time.Now().Add(time.Minute * 30).Unix()
	authorizedRoles = []string{"TEACHER"}
	token           = ""
)

func (m *TestModel) Authorise(username, password string) error {
	for _, m := range testModels {
		if m.User.Username == username && m.User.Password == password {
			return nil
		}
	}
	return errors.New("Unauthorised")
}

func init() {
	auth.SetPrivateKey(privateKey)
	auth.SetContextKey(contextClaimName)
	auth.SetRoleKeyName(roleKeyName)

	for _, ro := range testRoutes {
		handler := http.HandlerFunc(ro.handler)

		if len(ro.roles) != 0 {
			handler = auth.Scope(ro.roles, handler).(http.HandlerFunc)
		}

		if ro.auth {
			handler = auth.Validate(handler).(http.HandlerFunc)
		}
		r.Handle(ro.path, handler)
	}
}

func TestMain(t *testing.T) {
	t.Run("test 1st route", testRoute1)
	t.Run("test 2nd route", testRoute2)
}

var testRoute2 = func(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/basic/", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.OK(t, err)
	assert.Equal(string(body), "secret message", t)
}

type myClaims struct {
	Username string
	Role     string
	jwt.StandardClaims
}

var testRoute1 = func(t *testing.T) {
	for _, m := range testModels {
		t.Run(m.Name, func(t *testing.T) {
			t.Run("login", func(t *testing.T) {
				loginWriter := httptest.NewRecorder()
				formString := fmt.Sprintf(`{"Username":"%s", "Password":"%s"}`, m.User.Username, m.User.Password)
				postForm := strings.NewReader(formString)
				loginReq := httptest.NewRequest("POST", "/login/", postForm)
				loginReq.Header.Set("Content-Type", "application/json")
				r.ServeHTTP(loginWriter, loginReq)

				loginResp := loginWriter.Result()
				tokenBytes, err := ioutil.ReadAll(loginResp.Body)
				assert.OK(t, err)
				token = string(tokenBytes)
			})

			t.Run("test 3rd route", func(t *testing.T) {
				w := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/auth/", nil)
				req.Header.Set(roleKeyName, token)
				r.ServeHTTP(w, req)

				resp := w.Result()
				body, err := ioutil.ReadAll(resp.Body)
				assert.OK(t, err)

				if m.User.Password == m.Credential.Password && in(m.User.Role, authorizedRoles) {
					assert.Equal(string(body), "secret message", t)
				} else {
					assert.Panic(m.Name, t, func() {
						panic(string(body))
					})
				}
			})
		})
	}
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("secret message"))
}
func in(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	type LoginForm struct {
		Username string
		Password string
	}
	errCode := http.StatusUnauthorized
	loginForm := new(LoginForm)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.PrintError(w, err, errCode)
		return
	}
	if err := json.Unmarshal(body, loginForm); err != nil {
		helper.PrintError(w, err, errCode)
		return
	}

	username := loginForm.Username
	password := loginForm.Password

	for _, m := range testModels {
		if m.User.Username == username {
			claim := myClaims{
				Username: m.User.Username,
				Role:     m.User.Role,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expireToken,
				},
			}
			token, err := auth.CreateToken(claim, m, m.User.Username, password)
			if err != nil {
				helper.PrintError(w, err, errCode)
				return
			}
			w.Write([]byte(token))
			return
		}
	}
	helper.PrintError(w, errors.New("User not found."), errCode)
}
