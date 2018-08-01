package student_test

import (
	"fmt"
	"testing"

	"github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/model"
	"golang.org/x/crypto/bcrypt"
)

// create a student list for testing
var studentList = []model.Student{
	model.Student{
		User: model.User{
			Username: "lpstudent1",
			Password: "password1",
			Name:     "Alice Li",
			Cname:    "李麗絲",
		},
		ClassCode:   "3A",
		ClassNo:     1,
		Priority:    []int{0, 1, 2, 3},
		IsConfirmed: false,
		Rank:        -1,
	},
	model.Student{
		User: model.User{
			Username: "lpstudent2",
			Password: "password2",
			Name:     "Bob Li",
			Cname:    "李鮑伯",
		},
		ClassCode:   "3A",
		ClassNo:     2,
		Priority:    []int{3, 2, 1, 0},
		IsConfirmed: false,
		Rank:        -1,
	},
	model.Student{
		User: model.User{
			Username: "lpstudent3",
			Password: "password3",
			Name:     "Charlie Li",
			Cname:    "李查利",
		},
		ClassCode:   "3A",
		ClassNo:     3,
		Priority:    []int{},
		IsConfirmed: true,
		Rank:        -1},
}

func TestStudent(t *testing.T) {
	for i, sts := range studentList {
		name := fmt.Sprintf("Insert %d", i)
		t.Run(name, func(t *testing.T) {
			assert.OK(t, repo.StudentDB.Insert(&sts))
		})
	}

	t.Run("Insert Duplicated student", func(t *testing.T) {
		assert.Panic("Insert Duplicated student", t, func() {
			s := studentList[0]
			if err := repo.StudentDB.Insert(&s); err != nil {
				panic(err)
			}
		})
	})

	t.Run("studentList_get", func(t *testing.T) {
		students, err := repo.StudentDB.List()
		assert.OK(t, err)

		for i, got := range students {
			want := &studentList[i]

			// reset hashedPassword to un-hashed one for check the diff
			got.Password = want.Password
			assert.Equal(got, want, t)
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
			assert.OK(t, repo.StudentDB.UpdatePriority(student.Username, newPriority))
			// update the values in the student list for later checking
			studentList[i].Priority = newPriority
		})
	}

	for i, s := range studentList {
		name := fmt.Sprintf("GET %d", i)
		t.Run(name, func(t *testing.T) {
			username := s.Username
			got, err := repo.StudentDB.Get(username)
			assert.OK(t, err)

			want := &studentList[i]

			assert.OK(t, bcrypt.CompareHashAndPassword(
				[]byte(got.Password),
				[]byte(want.Password),
			))

			got.Password = want.Password
			assert.Equal(got, want, t)
		})
	}
}
