package main

import (
	"fmt"

	"./model"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := model.InitDB("./database/test.db")
	if err != nil {
		panic(err)
	}

	if err = db.CreateStudentTable(); err != nil {
		panic(err)
	}

	studentList := []model.Student{
		model.Student{"cyrusn", "顏昭洋"},
		model.Student{"winnie", "陳欣兒"},
	}
	for _, sts := range studentList {
		if err = db.AddStudent(sts); err != nil {
			panic(err)
		}
	}

	students, err := db.AllStudents()
	if err != nil {
		panic(err)
	}

	for _, sts := range students {
		fmt.Println(sts)
	}
}
