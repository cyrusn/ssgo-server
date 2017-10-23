package model_test

import (
	"fmt"
	"log"

	"github.com/cyrusn/ssgo/model"
)

func ExampleInitDB() {
	path := "./testing.db"
	db, err := model.InitDB(path)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleInsert() {
	model.InitDB("./testing.db")

	u := model.Student{
		user.Info{"lpstudent1", "password1", "Alice Li", "李麗絲"},
		"3A", 1, []int{0, 1, 2, 3}, false, -1,
	}

	if err := u.Insert(); err != nil {
		log.Fatal(err)
	}
	// Successful insert if err is nil
}

func ExampleGet() {
	db, err := ssdb.InitDB("./testing.db")
	if err != nil {
		log.Fatal(err)
	}

	u, err := student.Get(db, "student1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Name)
}
