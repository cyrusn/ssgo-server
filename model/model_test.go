package model_test

import (
	"log"
	"os"
	"testing"

	"github.com/cyrusn/ssgo/model"
	_ "github.com/mattn/go-sqlite3"

	"github.com/cyrusn/goTestHelper"
)

const (
	DBPath = "../database/test.db"
)

var repo *model.Repository

func init() {
	log.SetFlags(log.LstdFlags + log.Lshortfile)
}

func TestMain(m *testing.M) {
	log.Println(`Cleaning up DB: `, DBPath)
	cleanup(DBPath)
	repo = model.NewRepository(DBPath)
	os.Exit(m.Run())
}

// cleanup remove DB if it is exist and create a new empty database
func cleanup(DBPath string) {
	if _, err := os.Stat(DBPath); os.IsExist(err) {
		if err := os.Remove(DBPath); err != nil {
			log.Fatal(err)
		}
	}
	os.Create(DBPath)
}

func TestModel(t *testing.T) {
	t.Run("[Error] Init DB with invalid path", func(t *testing.T) {
		assert.Panic(`InitDB with an invalid path e.g. "./"`, t, func() {
			repo = model.NewRepository("./")
		})
	})
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
