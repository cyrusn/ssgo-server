package handlers

// httpLogger is a simple logger to log the request info
import (
	"encoding/json"
	"log"
	"net/http"
)

func httpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func jsonPrint(w http.ResponseWriter, v interface{}, errCode int) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		errPrint(w, err, errCode)
	}
}

func errPrint(w http.ResponseWriter, err error, errCode int) {
	http.Error(w, err.Error(), errCode)
}
