package model

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DBPath = "../database/test.db"
)

var db *DB

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

func TestCreateStudentTable(t *testing.T) {
	if err := db.CreateStudentTable(); err != nil {
		panic(err)
	}
}

func TestAddStudent(t *testing.T) {
	studentList := []Student{
		Student{"lpcyn", "thisIsPassword1", "cyrusn", "顏昭洋", []int{1, 3, 0, 2}, false},
		Student{"winnicC", "thisIsPassword1", "winnie", "陳欣兒", []int{3, 1, 2, 0}, true},
		Student{"alvinN", "thisIsPassword3", "alvin", "顏列藝", []int{}, true},
	}
	for _, sts := range studentList {
		if err := db.AddStudent(sts); err != nil {
			panic(err)
		}
	}
}

func TestAllStudents(t *testing.T) {
	students, err := db.AllStudents()
	if err != nil {
		panic(err)
	}

	for _, sts := range students {
		fmt.Println(sts)
	}
}
