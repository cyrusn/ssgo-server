package model_test

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/cyrusn/ssgo/model"
	_ "github.com/mattn/go-sqlite3"
)

var teachers = []model.Teacher{
	model.Teacher{"lpteacher1", "password1", "Alice Li", "李麗絲"},
	model.Teacher{"lpteacher2", "password2", "Bob Li", "李鮑伯"},
}

func TestTeacherPackage(t *testing.T) {
	for i, teacher := range teachers {
		name := fmt.Sprintf("Teacher_Insert %d", i)
		t.Run(name, func(t *testing.T) {
			if err := teacher.Insert(); err != nil {
				t.Fatal(err)
			}
		})
	}

	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("Teacher_Get %d", i)
		t.Run(name, func(t *testing.T) {
			want := teachers[i]
			teacher := new(model.Teacher)
			teacher.Username = want.Username
			err := teacher.Get()
			if err != nil {
				t.Fatal(err)
			}

			if err := bcrypt.CompareHashAndPassword(
				[]byte(teacher.Password),
				[]byte(want.Password),
			); err != nil {
				t.Fatal(err)
			}

			teacher.Password = want.Password
			diffTest(teacher, &want, t)
		})
	}
}
