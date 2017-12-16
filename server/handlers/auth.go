package handlers

import "net/http"

//
func (env *Env) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// check login
	// add jwt token to header

}

func (env *Env) IsAuth(r *http.Request) error { return nil }
