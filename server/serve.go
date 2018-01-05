package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	helper "github.com/cyrusn/goHTTPHelper"
	auth "github.com/cyrusn/goJWTAuthHelper"
	"github.com/cyrusn/ssgo/server/handlers"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

// Serve serve the routers
func Serve(env *handlers.Env) {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.

	for _, route := range handlers.Routes(env) {
		r1 := r.Methods(route.Methods...).
			Path(route.Path)

		if route.Auth {
			handler := http.HandlerFunc(route.Handler)
			r1 = r1.Handler(auth.Validate(handler))
		} else {
			r1 = r1.HandlerFunc(route.Handler)
		}

	}

	srv := &http.Server{
		Handler: helper.Logger(r),
		Addr:    "localhost" + env.Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Bind to a port and pass our router in
	fmt.Println("Available on http://localhost" + env.Port)
	log.Fatal(srv.ListenAndServe())
}
