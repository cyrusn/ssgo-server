package main

import (
	"github.com/cyrusn/ssgo/cli"
	"github.com/cyrusn/ssgo/model"
	"github.com/cyrusn/ssgo/server"
	"github.com/cyrusn/ssgo/server/handlers"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	cli.Parse()
}

func main() {
	const DBPath = "./database/test.db"
	var env handlers.Env
	var repo *model.Repository
	env.Port = cli.Port
	env.Vars = mux.Vars
	repo = model.NewRepository(DBPath)
	env.StudentStore = repo.StudentDB
	server.Serve(env)
}
