package teacher_test

import (
	"log"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/cyrusn/ssgo/model/helper"
	"github.com/cyrusn/ssgo/model/ssdb"
	"github.com/cyrusn/ssgo/model/teacher"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DBPath = "../../database/test.db"
)

var db *ssdb.DB

func init() {
	ssdb.Cleanup(DBPath)
	var err error

	db, err = ssdb.InitDB(DBPath)
	if err != nil {
		log.Fatal(err)
	}
	db.CreateTables()
}

var teachers = []teacher.Teacher{
	teacher.Teacher{"lpteacher1", "password1", "Alice Li", "李麗絲"},
	teacher.Teacher{"lpteacher2", "password2", "Bob Li", "李鮑伯"},
}

func Test(t *testing.T) {
	t.Run("Insert teacher user", TestInsert)
	t.Run("Insert teacher user", TestGet(0))
}

var TestInsert = func(t *testing.T) {
	for _, tr := range teachers {
		teacher.Insert(db, tr)
	}
}

var TestGet = func(index int) func(*testing.T) {
	return func(t *testing.T) {

		want := teachers[index]
		username := want.Username
		got, err := teacher.Get(db, username)
		if err != nil {
			t.Fatal(err)
		}

		if err := bcrypt.CompareHashAndPassword(
			[]byte(got.Password),
			[]byte(want.Password),
		); err != nil {
			t.Fatal(err)
		}

		got.Password = want.Password

		helper.DiffTest(got, &want, t)
	}
}
