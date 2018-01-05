package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cyrusn/goHTTPHelper"
	auth "github.com/cyrusn/goJWTAuthHelper"
	jwt "github.com/dgrijalva/jwt-go"
)

// Authenticator is an interface for Authentication
type Authenticator interface {
	Authenticate(loginName, password string) error
}

var (
	expireTime = time.Minute * time.Duration(30)
)

func expireToken(expireTime time.Duration) int64 {
	return time.Now().Add(expireTime).Unix()
}

// SetExpireTime set the expireTime for jwt
func SetExpireTime(duration time.Duration) {
	expireTime = duration
}

type authClaims struct {
	Username string
	Role     string
	jwt.StandardClaims
}

// StudentLoginHandler get jwt token for "STUDENT"
func (env *Env) StudentLoginHandler(w http.ResponseWriter, r *http.Request) {
	loginHandlerBuilder("STUDENT", env.StudentStore, w, r)
}

// TeacherLoginHandler get jwt token for "TEACHER"
func (env *Env) TeacherLoginHandler(w http.ResponseWriter, r *http.Request) {
	loginHandlerBuilder("TEACHER", env.TeacherStore, w, r)
}

// RefreshHandler refreshs the token
func (env *Env) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	errCode := http.StatusForbidden
	token := r.Header.Get(auth.GetJWTKeyName())

	encodedPayload := strings.Split(token, ".")[1]

	payload, err := jwt.DecodeSegment(encodedPayload)
	if err != nil {
		helper.PrintError(w, err, errCode)
	}
	claims := new(authClaims)

	err = json.Unmarshal(payload, claims)
	if err != nil {
		helper.PrintError(w, err, errCode)
	}

	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(expireTime).Unix()

	newToken, err := auth.CreateToken(claims)
	if err != nil {
		helper.PrintError(w, err, errCode)
	}
	w.Write([]byte(newToken))
}

func loginHandlerBuilder(role string, a Authenticator, w http.ResponseWriter, r *http.Request) {
	username, password, err := parseJSONPostForm(r)
	if err != nil {
		errCode := http.StatusBadRequest
		helper.PrintError(w, err, errCode)
		return
	}

	claims := authClaims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken(expireTime),
			IssuedAt:  time.Now().Unix(),
		},
	}

	if err = a.Authenticate(username, password); err != nil {
		errCode := http.StatusUnauthorized
		helper.PrintError(w, err, errCode)
		return
	}

	token, err := auth.CreateToken(claims)
	if err != nil {
		errCode := http.StatusUnauthorized
		helper.PrintError(w, err, errCode)
		return
	}

	w.Write([]byte(token))
}

func parseJSONPostForm(r *http.Request) (string, string, error) {
	type loginForm struct {
		Username string
		Password string
	}

	form := new(loginForm)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", "", err
	}

	if err = json.Unmarshal(body, form); err != nil {
		return "", "", err
	}

	return form.Username, form.Password, nil
}
