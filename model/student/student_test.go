package student_test

import (
	"log"
	"testing"

	"github.com/cyrusn/ssgo/model"
	"github.com/cyrusn/ssgo/model/helper"
	"github.com/cyrusn/ssgo/model/student"
	"github.com/cyrusn/ssgo/model/user"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DBPath = "../../database/test.db"
)

var db *model.DB

func init() {
	model.RemoveDB(DBPath)
	var err error

	db, err = model.InitDB(DBPath)
	if err != nil {
		log.Fatal(err)
	}
	db.CreateTables()
}

var studentList = []student.Student{
	student.Student{user.UserInfo{"lpstudent1", "password1", "Alice Li", "李愛麗"}, "3A", 1, []int{0, 1, 2, 3}, false, -1},
	student.Student{user.UserInfo{"lpstudent2", "password2", "Bob Li", "李鮑伯"}, "3A", 2, []int{3, 2, 1, 0}, false, -1},
	student.Student{user.UserInfo{"lpstudent3", "password3", "Charlie Li", "李查利"}, "3A", 3, []int{}, true, -1},
}

func Test(t *testing.T) {
	t.Run("Insert Students", TestInsertStudents)
	t.Run("List All Students", TestAllStudents)
	t.Run("Update student1 priority", TestUpdatePriorityInStudentsTable(0, []int{1, 2, 3, 0}))
	t.Run("Update student2 isConfirmed", TestUpdateIsConfirmedInStudentsTable(1, true))
	t.Run("List All Students", TestAllStudents)
	t.Run("Get student info", TestGetStudent(1))
}

var PanicTestAllStudent = func(t *testing.T) {
	helper.ExpectError("AllStudents", t, func() {
		if _, err := student.All(db); err != nil {
			panic(err)
		}
	})
}

var TestInsertStudents = func(t *testing.T) {
	for _, sts := range studentList {
		if err := student.Insert(db, sts); err != nil {
			t.Fatal(err)
		}
	}

	helper.ExpectError("InsertStudent", t, func() {
		s := student.Student{user.UserInfo{
			"lpstudent1",
			"password1",
			"Alice Li",
			"李愛麗",
		}, "3A", 1, []int{0, 1, 2, 3}, false, 2}
		if err := student.Insert(db, s); err != nil {
			panic(err)
		}
	})
}

var TestAllStudents = func(t *testing.T) {
	students, err := student.All(db)
	if err != nil {
		t.Fatal(err)
	}

	for i, got := range students {
		want := &studentList[i]
		helper.DiffTest(want, got, t)
	}
}

var TestUpdateIsConfirmedInStudentsTable = func(index int, newValue bool) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		if err := student.UpdateIsConfirmed(db, username, newValue); err != nil {
			t.Fatal(err)
		}
		studentList[index].IsConfirmed = newValue
	}
}

var TestUpdatePriorityInStudentsTable = func(index int, newPriority []int) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		err := student.UpdatePriority(db, username, newPriority)
		if err != nil {
			t.Fatal(err)
		}
		studentList[index].Priority = newPriority
	}
}

var TestGetStudent = func(index int) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		got, err := student.Get(db, username)
		if err != nil {
			t.Fatal(err)
		}

		want := &studentList[index]
		helper.DiffTest(want, got, t)
	}
}
