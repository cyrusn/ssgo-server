package model

import (
	"fmt"
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
	Student{"user2", "password2", "Bob", "鮑伯", []int{3, 2, 1, 0}, true},
	Student{"user3", "password3", "Charlie", "查利", []int{}, true},
}

func init() {
	if _, err := os.Stat(DBPath); os.IsExist(err) {
		os.Remove(DBPath)
	}
	os.Create(DBPath)
	var err error
	db, err = InitDB(DBPath)
	if err != nil {
		panic(err)
	}
}

func TestStudentPackage(t *testing.T) {
	t.Run("Create Database", TestCreateStudentTable)
	t.Run("Add Students", TestAddStudent)
	t.Run("List All Students", TestAllStudents)
	t.Run("Update user1 priority", TestUpdatePriorityInStudentsTable)
	t.Run("List All Students", TestAllStudents)
}

var TestCreateStudentTable = func(t *testing.T) {
	if err := db.CreateStudentTable(); err != nil {
		panic(err)
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

	for i, sts := range students {
		dummy := studentList[i]
		s := reflect.ValueOf(&dummy).Elem()
		st := reflect.ValueOf(sts).Elem()

		for j := 0; j < s.NumField(); j++ {
			want := st.Field(j).Interface()
			got := s.Field(j).Interface()
			fieldName := s.Type().Field(j).Name

			switch v := got.(type) {
			case string, bool:
				if want != got {
					t.Error(errMsg(fieldName, got, want))
				}
			case []int:
				if !reflect.DeepEqual(v, want.([]int)) {
					t.Error(errMsg(fieldName, got, want))
				}
			}
		}
	}
}

func errMsg(fieldName string, got, want interface{}) string {
	return fmt.Sprintf(
		"%s was incorrect, got: %v, want: %v.\n",
		fieldName,
		got,
		want,
	)
}

var TestUpdatePriorityInStudentsTable = func(t *testing.T) {
	newPriority := []int{1, 2, 3, 0}
	err := db.UpdatePriorityInStudentsTable("user1", newPriority)
	if err != nil {
		panic(err)
	}
	studentList[0].Priority = newPriority
}
