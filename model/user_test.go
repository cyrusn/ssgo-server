package model_test

import (
	"testing"

	"github.com/cyrusn/ssgo/model"
)

var userList = []model.User{
	model.User{"student1", "password1", "Alice", "愛麗絲", false},
	model.User{"student2", "password2", "Bob", "鮑伯", false},
	model.User{"student3", "password3", "Charlie", "查利", false},
	model.User{"teacher1", "password4", "Dave", "戴夫", true},
	model.User{"teacher2", "password5", "Eve", "伊夫", true},
	model.User{"teacher3", "password6", "Frank", "佛蘭克", true},
}

var TestUserTable = func(t *testing.T) {
	t.Run("Create user table", TestCreateUserTable)
	t.Run("Add users", TestInsertUser)
	t.Run("List All user", TestAllUsers)
	t.Run("Get user info", TestGetUser(1))
}

var TestCreateUserTable = func(t *testing.T) {
	if err := db.CreateUserTable(); err != nil {
		t.Fatal(err)
	}
}

var TestInsertUser = func(t *testing.T) {
	for _, u := range userList {
		if err := db.InsertUser(u); err != nil {
			t.Fatal(err)
		}
	}
}

var TestAllUsers = func(t *testing.T) {
	users, err := db.AllUsers()
	if err != nil {
		t.Fatal(err)
	}

	for i, got := range users {
		want := &userList[i]
		diffTest(want, got, t)
	}
}

var TestGetUser = func(index int) func(*testing.T) {
	return func(t *testing.T) {
		username := userList[index].Username
		got, err := db.GetUser(username)
		if err != nil {
			t.Error(err)
		}
		want := &userList[index]
		diffTest(want, got, t)
	}
}
