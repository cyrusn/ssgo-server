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

var studentList = []Student{
	Student{"user1", "password1", "Alice", "愛麗絲", []int{0, 1, 2, 3}, false},
	Student{"user2", "password2", "Bob", "鮑伯", []int{3, 2, 1, 0}, false},
	Student{"user3", "password3", "Charlie", "查利", []int{}, true},
}

var studentRanks = []Rank{
	Rank{"user1", 3},
	Rank{"user2", 1},
	Rank{"user3", 2},
}

func init() {
	cleanUpDB(DBPath)

	var err error
	db, err = InitDB(DBPath)
	if err != nil {
		panic(err)
	}
}

func TestStudentTable(t *testing.T) {
	t.Run("Create student table", TestCreateStudentTable)
	t.Run("Add Students", TestAddStudent)
	t.Run("List All Students", TestAllStudents)
	t.Run("Update user1 priority", TestUpdatePriorityInStudentsTable)
	t.Run("Update user2 isConfirmed", TestUpdateIsConfirmedInStudentsTable)
	t.Run("List All Students", TestAllStudents)
	t.Run("Get user3 info", TestGetStudent)
}

func TestRankTable(t *testing.T) {
	t.Run("Create rank table", TestCreateRankTable)
	t.Run("Insert rank data", TestInsertRankTable)
	t.Run("Get user2 rank", TestGetStudentRank)
}

var TestGetStudentRank = func(t *testing.T) {
	rank, err := db.GetStudentRank("user2")
	if err != nil {
		t.Error(err)
	}

	want := &studentRanks[1]
	got := rank

	diffTest(want, got, t)
}

var TestInsertRankTable = func(t *testing.T) {
	for _, rank := range studentRanks {
		if err := db.InsertStudentRank(rank); err != nil {
			t.Error(err)
		}
	}
}

var TestCreateRankTable = func(t *testing.T) {
	if err := db.CreateRankTable(); err != nil {
		t.Error(err)
	}
}

var TestCreateStudentTable = func(t *testing.T) {
	if err := db.CreateStudentTable(); err != nil {
		t.Error(err)
	}
}

var TestAddStudent = func(t *testing.T) {
	for _, sts := range studentList {
		if err := db.AddStudent(sts); err != nil {
			t.Error(err)
		}
	}
}

var TestAllStudents = func(t *testing.T) {
	students, err := db.AllStudents()
	if err != nil {
		t.Error(err)
	}

	for i, got := range students {
		want := &studentList[i]
		diffTest(want, got, t)
	}
}

var TestUpdateIsConfirmedInStudentsTable = func(t *testing.T) {
	newValue := true
	if err := db.UpdateIsConfirmedInStudentsTable("user2", newValue); err != nil {
		t.Error(err)
	}
	studentList[1].IsConfirmed = newValue
}

var TestUpdatePriorityInStudentsTable = func(t *testing.T) {
	newPriority := []int{1, 2, 3, 0}
	err := db.UpdatePriorityInStudentsTable("user1", newPriority)
	if err != nil {
		t.Error(err)
	}
	studentList[0].Priority = newPriority
}

var TestGetStudent = func(t *testing.T) {
	got, err := db.GetStudent("user3")
	if err != nil {
		t.Error(err)
	}

	want := &studentList[2]
	diffTest(want, got, t)
}

func diffTest(want, got interface{}, t *testing.T) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			want,
			got,
		)
	}
}

func cleanUpDB(DBPath string) {
	if _, err := os.Stat(DBPath); os.IsExist(err) {
		os.Remove(DBPath)
	}
	os.Create(DBPath)
}
