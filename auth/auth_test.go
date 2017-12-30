package auth_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/auth"
	jwt "github.com/dgrijalva/jwt-go"
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

func (m *TestModel) Authorise(username, password string) error {
	for _, m := range testModels {
		if m.User.Username == username && m.User.Password == password {
			return nil
		}
	}
	return errors.New("Unauthorised")
}

type payload struct {
	Username string
	Role     string
}

const (
	AuthorizedRole = "TEACHER"
)

var (
	privateKey = []byte("hello world")
	tokens     = []string{}
)

func init() {
	auth.SetPrivateKey(privateKey)
}

func TestMain(t *testing.T) {
	expireToken := time.Now().Add(time.Minute * 30).Unix()
	handler := http.HandlerFunc(simpleHandler)
	for _, m := range testModels {
		t.Run(m.Name, func(t *testing.T) {

			claim := auth.Claim{
				Payload: payload{Username: m.User.Username, Role: m.User.Role},
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expireToken,
				},
			}
			password := m.Credential.Password
			token, err := claim.GetToken(m, m.User.Username, password)

			if m.User.Password == password {
				assert.OK(t, err)
			} else {
				assert.Panic("Fail to login", t, func() {
					panic(err)
				})
			}
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/auth/", nil)
			req.Header.Set("jwt-token", token)
			authorizedHandler := auth.Required(handler)
			authorizedHandler.ServeHTTP(w, req)
			resp := w.Result()
			body, err := ioutil.ReadAll(resp.Body)
			assert.OK(t, err)

			if m.User.Role == AuthorizedRole {
				assert.Equal(string(body), "secret message", t)
			} else {
				assert.NotEqual(string(body), "secret message", t)
			}
		})
	}
}

func parseRole(ctx context.Context) (string, error) {
	m := ctx.Value(auth.ContextKeyClaim).(jwt.MapClaims)
	payload, ok := m["Payload"].(map[string]interface{})
	if !ok {
		return "", errors.New("Claim not found in context")
	}

	result, ok := payload["Role"].(string)
	if !ok {
		return "", errors.New("Invalid payload")
	}
	return result, nil
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	errCode := http.StatusUnauthorized
	role, err := parseRole(r.Context())
	if err != nil {
		helper.PrintError(w, err, errCode)
	}
	if role == AuthorizedRole {
		w.Write([]byte("secret message"))
		return
	}
	helper.PrintError(w, errors.New("User not allow"), errCode)
}
