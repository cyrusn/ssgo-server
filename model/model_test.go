package model_test

import (
	"testing"

	"github.com/cyrusn/ssgo/model"
	_ "github.com/mattn/go-sqlite3"
)

var PanicTestInitDB = func(t *testing.T) {
	expectError(`InitDB with an invalid path e.g. "./"`, t, func() {
		model.InitDB("./")
	})
}

var PanicTestCreateTables = func(t *testing.T) {
	expectError("CreateTables before DB ready", t, func() {
		if err := model.CreateTables(); err != nil {
			panic(err)
		}
	})
}

var TestInitDB = func(DBPath string) func(*testing.T) {
	return func(t *testing.T) {
		model.InitDB(DBPath)
	}
}

var TestCreateTable = func(t *testing.T) {
	err := model.CreateTables()
	if err != nil {
		t.Fatal(err)
	}
}
