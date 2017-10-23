package model_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// create a student list for testing
var studentList = []model.Student{
	model.Student{user.Info{"lpstudent1", "password1", "Alice Li", "李麗絲"}, "3A", 1, []int{0, 1, 2, 3}, false, -1},
	model.Student{user.Info{"lpstudent2", "password2", "Bob Li", "李鮑伯"}, "3A", 2, []int{3, 2, 1, 0}, false, -1},
	model.Student{user.Info{"lpstudent3", "password3", "Charlie Li", "李查利"}, "3A", 3, []int{}, true, -1},
}

func TestStudent_Insert(t *testing.T) {
	for _, sts := range studentList {
		if err := student.Insert(); err != nil {
			t.Fatal(err)
		}
	}
}

// TestInsertError generate error for Insert function, to test the error handling
func TestStudent_Insert_Errors(t *testing.T) {
	expectError("Insert Duplicated student", t, func() {
		s := studentList[0]
		if err := student.Insert(db, s); err != nil {
			panic(err)
		}
	})
}

func TestStudent_All(t *testing.T) {
	var students []*model.Student
	err := students.All()
	if err != nil {
		t.Fatal(err)
	}

	for i, got := range students {
		want := &studentList[i]

		// reset hashedPassword to un-hashed one for check the diff
		got.Info.Password = want.Info.Password

		diffTest(got, want, t)
	}
}

func TestStudent_UpdateIsConfirmed(*testing.T) {
	for i := 0; i < 2; i++ {
		username := studentList[index].Username
		var student model.Student
		student.Username = username
		if err := student.UpdateIsConfirmed(true); err != nil {
			t.Fatal(err)
		}
		// update the values in the student list for later checking
		studentList[index].IsConfirmed = newValue
		diffTest(student, want, t)
	}
}

func TestStudent_UpdatePriority(*testing.T) {
	newPriority := []int{1, 2, 3, 0}
	username := studentList[0].Username
	var student = new(model.Student)
	student.Username = username
	err := student.UpdatePriority(newPriority)
	if err != nil {
		t.Fatal(err)
	}
	// update the values in the student list for later checking
	studentList[index].Priority = newPriority
	diffTest(student, want, t)
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
