package model_test

import (
	"testing"

	"github.com/cyrusn/ssgo/model"
)

var studentList = []model.Student{
	model.Student{"student1", "3A", 1, []int{0, 1, 2, 3}, false},
	model.Student{"student2", "3A", 2, []int{3, 2, 1, 0}, false},
	model.Student{"student3", "3A", 3, []int{}, true},
}

var TestStudentTable = func(t *testing.T) {
	t.Run("Panic Test All Students", PanicTestAllStudent)
	t.Run("Create student table", TestCreateStudentTable)
	t.Run("Add Students", TestInsertStudent)
	t.Run("List All Students", TestAllStudents)
	t.Run("Update student1 priority", TestUpdatePriorityInStudentsTable(0, []int{1, 2, 3, 0}))
	t.Run("Update student2 isConfirmed", TestUpdateIsConfirmedInStudentsTable(1, true))
	t.Run("List All Students", TestAllStudents)
	t.Run("Get student info", TestGetStudent(1))
}

var PanicTestAllStudent = func(t *testing.T) {
	expectError("AllStudents", t, func() {
		if _, err := db.AllStudents(); err != nil {
			panic(err)
		}
	})
}

var TestCreateStudentTable = func(t *testing.T) {
	if err := db.CreateStudentTable(); err != nil {
		t.Fatal(err)
	}
}

var TestInsertStudent = func(t *testing.T) {
	for _, sts := range studentList {
		if err := db.InsertStudent(sts); err != nil {
			t.Fatal(err)
		}
	}

	expectError("InsertStudent", t, func() {
		s := model.Student{"student1", "3A", 1, []int{0, 1, 2, 3}, false}
		if err := db.InsertStudent(s); err != nil {
			panic(err)
		}
	})
}

var TestAllStudents = func(t *testing.T) {
	students, err := db.AllStudents()
	if err != nil {
		t.Fatal(err)
	}

	for i, got := range students {
		want := &studentList[i]
		diffTest(want, got, t)
	}
}

var TestUpdateIsConfirmedInStudentsTable = func(index int, newValue bool) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		if err := db.UpdateIsConfirmedInStudentsTable(username, newValue); err != nil {
			t.Fatal(err)
		}
		studentList[index].IsConfirmed = newValue
	}
}

var TestUpdatePriorityInStudentsTable = func(index int, newPriority []int) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		err := db.UpdatePriorityInStudentsTable(username, newPriority)
		if err != nil {
			t.Fatal(err)
		}
		studentList[index].Priority = newPriority
	}
}

var TestGetStudent = func(index int) func(*testing.T) {
	return func(t *testing.T) {
		username := studentList[index].Username
		got, err := db.GetStudent(username)
		if err != nil {
			t.Fatal(err)
		}

		want := &studentList[index]
		diffTest(want, got, t)
	}

}
