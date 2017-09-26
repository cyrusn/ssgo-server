package model_test

import (
	"testing"

	"github.com/cyrusn/ssgo/model"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBPath = "../database/test.db"
)

var db *model.DB

func init() {
	cleanUpDB(DBPath)
}

func Test(t *testing.T) {
	t.Run("[Panic] CreateTables", PanicTestCreateTables)
	t.Run("[Panic] InitDB", PanicTestInitDB)
	t.Run("InitDB", TestInitDB(DBPath))
	t.Run("CreateTables", TestCreateTable)
	// t.Run("User", TestUserTable)
}

var PanicTestInitDB = func(t *testing.T) {
	expectError(`InitDB with an invalid path e.g. "./"`, t, func() {
		if _, err := model.InitDB("./"); err != nil {
			panic(err)
		}
	})
}

var TestInitDB = func(DBPath string) func(*testing.T) {
	return func(t *testing.T) {
		var err error
		db, err = model.InitDB(DBPath)
		if err != nil {
			t.Fatal(err)
		}
	}
}
