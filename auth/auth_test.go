package auth_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/auth"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const roleKeyName = "kid"

var (
	r               = mux.NewRouter()
	expireToken     = time.Now().Add(time.Minute * 30).Unix()
	privateKey      = []byte("hello world")
	tokens          = []string{}
	authorizedRoles = []string{"TEACHER"}
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
	auth.SetContextKey("myContextClaim")
	auth.SetRoleKeyName(roleKeyName)

	for _, ro := range routes {
		handler := http.HandlerFunc(ro.handler)

		if len(ro.roles) != 0 {
			handler = auth.Scope(ro.roles, handler).(http.HandlerFunc)
		}

		if ro.auth {
			handler = auth.Required(handler).(http.HandlerFunc)
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
			token := ""
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
