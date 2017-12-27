package handlers_test

import (
	"github.com/cyrusn/ssgo/model"
	"github.com/cyrusn/ssgo/server/handlers"
	"github.com/gorilla/mux"
)

var env = &handlers.Env{
	StudentStore: store,
	Vars:         mux.Vars,
}

var studentList = []*model.Student{
	&model.Student{
		User:        model.User{Username: "lpstudent1", Password: "password1", Name: "Alice Li", Cname: "李麗絲"},
		ClassCode:   "3A",
		ClassNo:     1,
		Priority:    []int{0, 1, 2, 3},
		IsConfirmed: false,
		Rank:        -1,
	},
	&model.Student{
		User:        model.User{Username: "lpstudent2", Password: "password2", Name: "Bob Li", Cname: "李鮑伯"},
		ClassCode:   "3B",
		ClassNo:     2,
		Priority:    []int{3, 2, 1, 0},
		IsConfirmed: false,
		Rank:        -1,
	},
	&model.Student{
		User:        model.User{Username: "lpstudent3", Password: "password3", Name: "Charlie Li", Cname: "李查利"},
		ClassCode:   "3C",
		ClassNo:     3,
		Priority:    []int{},
		IsConfirmed: true,
		Rank:        -1,
	},
}
