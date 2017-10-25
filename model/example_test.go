package model_test

import (
	"fmt"
	"log"

	"github.com/cyrusn/ssgo/model"
)

func ExampleInitDB() {
	path := "./testing.db"
	model.InitDB(path)
}

func ExampleInsert() {
	model.InitDB("./testing.db")

	u := model.Student{
		model.Info{"lpstudent1", "password1", "Alice Li", "李麗絲"},
		"3A", 1, []int{0, 1, 2, 3}, false, -1,
	}

	if err := u.Insert(); err != nil {
		log.Fatal(err)
	}
	// Successful insert if err is nil
}

func ExampleGet() {
	model.InitDB("./testing.db")

	student := new(model.Student)
	student.Info.Username = "student1"
	err := student.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(student.Name)
}
