package main

import (
	"log"
	"net/http"

	"github.com/cyrusn/ssgo/model"
	"github.com/cyrusn/ssgo/server/handlers"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

var env handlers.Env
var repo *model.Repository

const (
	DBPath = "../database/test.db"
)

func init() {
	repo = model.NewRepository(DBPath)
	env.StudentStore = repo.StudentDB
	env.Vars = mux.Vars
}
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.

	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/students/{username}", env.GetStudentHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
