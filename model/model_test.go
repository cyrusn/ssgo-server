package model_test

import (
	"os"
	"reflect"
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
	t.Run("[Panic Test] CreateTable", PanicTestCreateTable)
	t.Run("[Panic Test] InitDB", PanicTestInitDB)
	t.Run("InitDB", TestInitDB(DBPath))
	t.Run("User", TestUserTable)
	t.Run("Student", TestStudentTable)
	t.Run("Rank", TestRankTable)
	t.Run("Subject", TestSubjectTable)
}

var PanicTestInitDB = func(t *testing.T) {
	expectError("InitDB", t, func() {
		if _, err := model.InitDB("./"); err != nil {
			panic(err)
		}
	})
}

var PanicTestCreateTable = func(t *testing.T) {
	type testTable struct {
		name   string
		method func() error
	}

	var testTables = []testTable{
		testTable{"CreateRankTable", db.CreateRankTable},
		testTable{"CreateStudentTable", db.CreateStudentTable},
		testTable{"CreateSubjectTable", db.CreateSubjectTable},
		testTable{"CreateUserTable", db.CreateUserTable},
	}

	for _, table := range testTables {
		expectError(table.name, t, func() {
			err := table.method()
			if err != nil {
				panic(err)
			}
		})
	}

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

func expectError(name string, t *testing.T, f func()) {
	defer func(t *testing.T) {
		err := recover()
		if err == nil {
			t.Fatalf("Error Test: [%s] did not return error", name)
		}
		t.Logf("Error Test: Success! [%s], %s", name, err)
	}(t)
	f()
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
