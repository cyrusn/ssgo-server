package model_test

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/cyrusn/ssgo/helper"
	"github.com/cyrusn/ssgo/model"
	_ "github.com/mattn/go-sqlite3"
)

var teachers = []model.Teacher{
	model.Teacher{"lpteacher1", "password1", "Alice Li", "李麗絲"},
	model.Teacher{"lpteacher2", "password2", "Bob Li", "李鮑伯"},
}

func TestTeacher(t *testing.T) {
	for i, teacher := range teachers {
		name := fmt.Sprintf("Teacher_Insert #%d", i+1)
		t.Run(name, func(t *testing.T) {
			if err := repo.TeacherDB.Insert(&teacher); err != nil {
				t.Fatal(err)
			}
		})
	}

	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("Teacher_Get #%d", i+1)
		t.Run(name, func(t *testing.T) {
			want := teachers[i]
			teacher, err := repo.TeacherDB.Get(want.Username)
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
			helper.DiffTest(teacher, &want, t)
		})
	}
}
