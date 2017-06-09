package model

import (
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DBPath = "../database/test.db"
)

var db *DB

func init() {
	cleanUpDB(DBPath)
}

func Test(t *testing.T) {
	t.Run("InitDB", TestInitDB(DBPath))
	t.Run("User", TestUserTable)
	t.Run("Student", TestStudentTable)
	t.Run("Rank", TestRankTable)
	t.Run("Subject", TestSubjectTable)
}

var TestInitDB = func(DBPath string) func(*testing.T) {
	return func(t *testing.T) {
		var err error
		db, err = InitDB(DBPath)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func diffTest(want, got interface{}, t *testing.T) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			got,
			want,
		)
	}
}

func cleanUpDB(DBPath string) {
	if _, err := os.Stat(DBPath); os.IsExist(err) {
		os.Remove(DBPath)
	}
	os.Create(DBPath)
}
