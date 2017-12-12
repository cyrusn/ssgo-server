package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cyrusn/ssgo/helper"
	"github.com/cyrusn/ssgo/server/handlers"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

func Serve(env handlers.Env) {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.

	for _, route := range Routes(env) {
		r.Methods(route.Methods...).
			Path(route.Path).
			HandlerFunc(route.Handler)
	}

	srv := &http.Server{
		Handler: helper.HTTPLogger(r),
		Addr:    "localhost" + env.Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Bind to a port and pass our router in
	fmt.Println("Available on http://localhost" + env.Port)
	log.Fatal(srv.ListenAndServe())
}
