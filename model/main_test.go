package model_test

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/cyrusn/ssgo/model"
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

// expectError is a testing tool, it used to test for error handling
func expectError(name string, t *testing.T, f func()) {
	defer func(t *testing.T) {
		err := recover()

		if err == nil {
			t.Fatalf("Error Test: [%s] did not return error", name)
		}
	}(t)
	f()
}

// diffTest is simply test if there are differences of 2 structs
func diffTest(got, want interface{}, t *testing.T) {
	if !reflect.DeepEqual(want, got) {

		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			got,
			want,
		)
	}
}
