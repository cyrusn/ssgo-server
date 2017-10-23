package model_test

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

var teachers = []teacher.Teacher{
	model.Teacher{"lpteacher1", "password1", "Alice Li", "李麗絲"},
	model.Teacher{"lpteacher2", "password2", "Bob Li", "李鮑伯"},
}

func TestTeacher_Insert(t *testing.T) {
	for _, tr := range teachers {
		teacher.Insert(db, tr)
	}
}

func TestTeacher_Get(*testing.T) {
	for i := 0; i < 2; i++ {
		want := teachers[i]
		teacher := new(model.Teacher)
		teacher.Username = want.Username
		err := teacher.Get()
		if err != nil {
			t.Fatal(err)
		}

		if err := bcrypt.CompareHashAndPassword(
			[]byte(got.Password),
			[]byte(want.Password),
		); err != nil {
			t.Fatal(err)
		}

		teacher.Password = want.Password
		diffTest(teacher, &want, t)
	}
}
