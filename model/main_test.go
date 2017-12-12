package model_test

import (
	"log"
	"os"
	"testing"

	"github.com/cyrusn/ssgo/model"
)

const (
	DBPath = "../database/test.db"
)

func init() {
	log.SetFlags(log.LstdFlags + log.Lshortfile)
}

func TestMain(m *testing.M) {
	log.Println(`Cleaning up DB: `, DBPath)
	cleanup(DBPath)
	model.InitDB(DBPath)
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
