package model_test

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/model"
	_ "github.com/mattn/go-sqlite3"
)

var teachers = []model.Teacher{
	model.Teacher{
		Username: "lpteacher1",
		Password: "password1",
		Name:     "Alice Li",
		Cname:    "李麗絲"},
	model.Teacher{
		Username: "lpteacher2",
		Password: "password2",
		Name:     "Bob Li",
		Cname:    "李鮑伯"},
}

func TestTeacher(t *testing.T) {
	for i, teacher := range teachers {
		name := fmt.Sprintf("Teacher_Insert #%d", i+1)
		t.Run(name, func(t *testing.T) {
			assert.OK(t, repo.TeacherDB.Insert(&teacher))
		})
	}

	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("Teacher_Get #%d", i+1)
		t.Run(name, func(t *testing.T) {
			want := teachers[i]
			teacher, err := repo.TeacherDB.Get(want.Username)
			assert.OK(t, err)

			assert.OK(t, bcrypt.CompareHashAndPassword(
				[]byte(teacher.Password),
				[]byte(want.Password),
			))
			teacher.Password = want.Password
			assert.Equal(teacher, &want, t)
		})
	}
}
