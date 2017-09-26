package student_test

import (
	"log"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/cyrusn/ssgo/model/helper"
	"github.com/cyrusn/ssgo/model/ssdb"
	"github.com/cyrusn/ssgo/model/student"

	"github.com/cyrusn/ssgo/model/user"

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

// create a student list for testing
var studentList = []student.Student{
	student.Student{user.Info{"lpstudent1", "password1", "Alice Li", "李麗絲"}, "3A", 1, []int{0, 1, 2, 3}, false, -1},
	student.Student{user.Info{"lpstudent2", "password2", "Bob Li", "李鮑伯"}, "3A", 2, []int{3, 2, 1, 0}, false, -1},
	student.Student{user.Info{"lpstudent3", "password3", "Charlie Li", "李查利"}, "3A", 3, []int{}, true, -1},
}

// Test is the main programme for testing, it runs the sub-test
func Test(t *testing.T) {
	t.Run("Insert all students", TestInsert)
	t.Run("Insert duplicated student", TestInsertError)
	t.Run("List all students", TestAll)
	t.Run("Get student", TestGet(0))
	t.Run("Update student1 priority", TestUpdatePriority(0, []int{1, 2, 3, 0}))
	t.Run("Update student1 isConfirmed", TestUpdateIsConfirmed(0, true))
	t.Run("Get student", TestGet(0))
}

var TestInsert = func(t *testing.T) {
	for _, sts := range studentList {
		if err := student.Insert(db, sts); err != nil {
			t.Fatal(err)
		}
	}
}

// TestInsertError generate error for Insert function, to test the error handling
var TestInsertError = func(t *testing.T) {
	helper.ExpectError("Insert Duplicated student", t, func() {
		s := studentList[0]
		if err := student.Insert(db, s); err != nil {
			panic(err)
		}
	})
}

var TestAll = func(t *testing.T) {
	students, err := student.All(db)
	if err != nil {
		t.Fatal(err)
	}

	for i, got := range students {
		want := &studentList[i]

		// reset hashedPassword to un-hashed one for check the diff
		got.Info.Password = want.Info.Password

		helper.DiffTest(got, want, t)

	}
}

var TestUpdateIsConfirmed = func(index int, newValue bool) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		if err := student.UpdateIsConfirmed(db, username, newValue); err != nil {
			t.Fatal(err)
		}
		// update the values in the student list for later checking
		studentList[index].IsConfirmed = newValue
	}
}

var TestUpdatePriority = func(index int, newPriority []int) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		err := student.UpdatePriority(db, username, newPriority)
		if err != nil {
			t.Fatal(err)
		}
		// update the values in the student list for later checking
		studentList[index].Priority = newPriority
	}
}

var TestGet = func(index int) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		got, err := student.Get(db, username)

		if err != nil {
			t.Fatal(err)
		}

		want := &studentList[index]

		if err := bcrypt.CompareHashAndPassword([]byte(got.Password), []byte(want.Password)); err != nil {
			t.Fatal(err)
		}

		got.Info.Password = want.Info.Password
		helper.DiffTest(got, want, t)
	}
}
