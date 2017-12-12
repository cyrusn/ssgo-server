package model_test

import (
	"testing"

	"github.com/cyrusn/ssgo/helper"
	"github.com/cyrusn/ssgo/model"
	_ "github.com/mattn/go-sqlite3"
)

func TestModel(t *testing.T) {
	t.Run("[Error] Init DB with invalid path", func(t *testing.T) {
		helper.ExpectError(`InitDB with an invalid path e.g. "./"`, t, func() {
			model.InitDB("./")
		})
	})

	t.Run("[Error] CreateTables without properly init DB", func(t *testing.T) {
		helper.ExpectError("CreateTables before DB ready", t, func() {
			if err := model.CreateTables(); err != nil {
				panic(err)
			}
		})
	})

	t.Run("Init DB", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Fatal(err)
			}
		}()
		model.InitDB(DBPath)
	})

	t.Run("CreateTables", func(t *testing.T) {
		err := model.CreateTables()
		if err != nil {
			t.Fatal(err)
		}
	})
}
