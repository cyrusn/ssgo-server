package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cyrusn/goHTTPHelper"
	auth "github.com/cyrusn/goJWTAuthHelper"
	jwt "github.com/dgrijalva/jwt-go"
)

type authClaim struct {
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

func loginHandlerBuilder(role string, a auth.Authenticator, w http.ResponseWriter, r *http.Request) {
	username, password, err := parseJSONPostForm(r)
	if err != nil {
		errCode := http.StatusBadRequest
		helper.PrintError(w, err, errCode)
		return
	}

	expireToken := time.Now().Add(time.Minute * 30).Unix()

	claim := authClaim{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}

	token, err := auth.CreateToken(claim, a, username, password)
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

// RefreshHandler refreshs the token
func (env *Env) RefreshHandler(w http.ResponseWriter, r *http.Request) {

}
