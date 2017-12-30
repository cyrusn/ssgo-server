package model_test

import (
	"fmt"
	"testing"

	"github.com/cyrusn/goTestHelper"
	"github.com/cyrusn/ssgo/model"
)

var subjectList = []model.Subject{
	model.Subject{
		Code:     "bio",
		Group:    1,
		Name:     "Biology",
		Cname:    "生物",
		Capacity: 0,
	},
	model.Subject{
		Code:     "bafs",
		Group:    1,
		Name:     "Business,Accounting and Financial Studies",
		Cname:    "企業、會計與財務概論",
		Capacity: 0,
	},
	model.Subject{
		Code:     "ict",
		Group:    2,
		Name:     "Information and Communication Technology",
		Cname:    "資訊及通訊科技",
		Capacity: 0,
	},
	model.Subject{
		Code:     "econ",
		Group:    2,
		Name:     "Economics",
		Cname:    "經濟",
		Capacity: 0,
	},
}

func TestSubject(t *testing.T) {
	for i, s := range subjectList {
		name := fmt.Sprintf("Insert Subject #%d", i+1)
		t.Run(name, func(t *testing.T) {
			assert.OK(t, repo.SubjectDB.Insert(&s))
		})
	}

	t.Run("Update Capacity", func(t *testing.T) {
		for i, subject := range subjectList {
			capacity := 20
			assert.OK(t, repo.SubjectDB.UpdateCapacity(subject.Code, capacity))
			subjectList[i].Capacity = capacity
		}
	})

	for i, subject := range subjectList {
		name := fmt.Sprintf("Get Each Subjects #%d", i+1)
		t.Run(name, func(t *testing.T) {
			want := subject
			subjectCode := want.Code
			s, err := repo.SubjectDB.Get(subjectCode)
			assert.OK(t, err)
			assert.Equal(&want, s, t)
		})
	}

	t.Run("Get All Subjects", func(t *testing.T) {
		subjects, err := repo.SubjectDB.List()
		assert.OK(t, err)
		for i, got := range subjects {
			assert.Equal(&subjectList[i], got, t)
		}
	})
}
