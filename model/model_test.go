package model_test

import (
	"testing"

	"github.com/cyrusn/ssgo/model"
	_ "github.com/mattn/go-sqlite3"

	"github.com/cyrusn/goTestHelper"
)

func TestModel(t *testing.T) {
	t.Run("[Error] Init DB with invalid path", func(t *testing.T) {
		assert.Panic(`InitDB with an invalid path e.g. "./"`, t, func() {
			repo = model.NewRepository("./")
		})
	})
	// t.Run("[Error] CreateTables without properly init DB", func(t *testing.T) {
	// 	assert.Panic("CreateTables before DB ready", t, func() {
	// 		if err := model.CreateTables(repo.DB); err != nil {
	// 			panic(err)
	// 		}
	// 	})
	// })
	t.Run("Init DB", func(t *testing.T) {
		defer func() {
			assert.OK(t, recover())
		}()
		repo = model.NewRepository(DBPath)
	})

	t.Run("CreateTables", func(t *testing.T) {
		assert.OK(t, model.CreateTables(repo.DB))
	})
}
