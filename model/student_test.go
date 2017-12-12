package model_test

import (
	"fmt"
	"testing"

	"github.com/cyrusn/ssgo/helper"
	"github.com/cyrusn/ssgo/model"
	"golang.org/x/crypto/bcrypt"
)

// create a student list for testing
var studentList = []model.Student{
	model.Student{model.User{"lpstudent1", "password1", "Alice Li", "李麗絲"}, "3A", 1, []int{0, 1, 2, 3}, false, -1},
	model.Student{model.User{"lpstudent2", "password2", "Bob Li", "李鮑伯"}, "3A", 2, []int{3, 2, 1, 0}, false, -1},
	model.Student{model.User{"lpstudent3", "password3", "Charlie Li", "李查利"}, "3A", 3, []int{}, true, -1},
}

func TestStudent(t *testing.T) {
	for i, sts := range studentList {
		name := fmt.Sprintf("Insert %d", i)
		t.Run(name, func(t *testing.T) {
			if err := repo.StudentDB.Insert(&sts); err != nil {
				t.Fatal(err)
			}
		})
	}

	t.Run("Insert Duplicated student", func(t *testing.T) {
		helper.ExpectError("Insert Duplicated student", t, func() {
			s := studentList[0]
			if err := repo.StudentDB.Insert(&s); err != nil {
				panic(err)
			}
		})
	})

	t.Run("studentList_get", func(t *testing.T) {
		students, err := repo.StudentDB.List()
		if err != nil {
			t.Fatal(err)
		}

		for i, got := range students {
			want := &studentList[i]

			// reset hashedPassword to un-hashed one for check the diff
			got.Password = want.Password
			helper.DiffTest(got, want, t)
		}
	})

	for i, student := range studentList {
		name := fmt.Sprintf("Update Is Confirmed #%d", i+1)
		t.Run(name, func(t *testing.T) {
			newValue := true
			if err := repo.StudentDB.UpdateIsConfirmed(student.Username, newValue); err != nil {
				t.Fatal(err)
			}
			// update the values in the student list for later checking
			studentList[i].IsConfirmed = newValue
			// helper.DiffTest(student, studentList[i], t)
		})
	}

	newPriorities := [][]int{
		[]int{1, 2, 3, 0},
		[]int{3, 2, 1, 0},
		[]int{1, 3, 0, 2},
	}

	for i, student := range studentList {
		name := fmt.Sprintf("Update Priority #%d", i+1)

		t.Run(name, func(t *testing.T) {
			newPriority := newPriorities[i]

			if err := repo.StudentDB.UpdatePriority(student.Username, newPriority); err != nil {
				t.Fatal(err)
			}
			// update the values in the student list for later checking
			studentList[i].Priority = newPriority
			// helper.DiffTest(s, studentList[0], t)
		})
	}

	for i, s := range studentList {
		name := fmt.Sprintf("GET %d", i)
		t.Run(name, func(t *testing.T) {
			username := s.Username
			got, err := repo.StudentDB.Get(username)
			if err != nil {
				t.Fatal(err)
			}

			want := &studentList[i]

			if err := bcrypt.CompareHashAndPassword([]byte(got.Password), []byte(want.Password)); err != nil {
				t.Fatal(err)
			}

			got.Password = want.Password
			helper.DiffTest(got, want, t)
		})
	}
}
