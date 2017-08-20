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

func ExampleDB_InsertUser() {
	db, err := model.InitDB("./testing.db")
	if err != nil {
		log.Fatal(err)
	}
	u := model.User{"student1", "password1", "Alice", "愛麗絲", false}
	if err := db.InsertUser(u); err != nil {
		log.Fatal(err)
	}
	// Successful insert if err is nil
}

func ExampleDB_GetUser() {
	db, err := model.InitDB("./testing.db")
	if err != nil {
		log.Fatal(err)
	}

	u, err := db.GetUser("student1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Name)
}

func ExampleUser_Validate() {
	db, err := model.InitDB("./testing.db")
	if err != nil {
		log.Fatal(err)
	}

	u, err := db.GetUser("student1")
	if err != nil {
		log.Fatal(err)
	}

	if err := u.Validate("password1"); err != nil {
		log.Fatal(err)
	}
}
