package model_test

import (
	"fmt"
	"testing"

	"github.com/cyrusn/ssgo/helper"
	"github.com/cyrusn/ssgo/model"
)

var subjectList = []model.Subject{
	model.Subject{"bio", 1, "Biology", "生物", 0},
	model.Subject{"bafs", 1, "Business, Accounting and Financial Studies", "企業、會計與財務概論", 0},
	model.Subject{"ict", 2, "Information and Communication Technology", "資訊及通訊科技", 0},
	model.Subject{"econ", 2, "Economics", "經濟", 0},
}

func TestSubject(t *testing.T) {
	for i, s := range subjectList {
		name := fmt.Sprintf("Insert Subject #%d", i+1)
		t.Run(name, func(t *testing.T) {
			if err := repo.SubjectDB.Insert(&s); err != nil {
				t.Fatal(err)
			}
		})
	}

	t.Run("Update Capacity", func(t *testing.T) {
		for i, subject := range subjectList {
			capacity := 20
			if err := repo.SubjectDB.UpdateCapacity(subject.Code, capacity); err != nil {
				t.Fatal(err)
			}
			subjectList[i].Capacity = capacity
		}
	})

	for i, subject := range subjectList {
		name := fmt.Sprintf("Get Each Subjects #%d", i+1)
		t.Run(name, func(t *testing.T) {
			want := subject
			subjectCode := want.Code
			s, err := repo.SubjectDB.Get(subjectCode)
			if err != nil {
				t.Fatal(err)
			}
			helper.DiffTest(&want, s, t)
		})
	}

	t.Run("Get All Subjects", func(t *testing.T) {
		subjects, err := repo.SubjectDB.List()
		if err != nil {
			t.Fatal(err)
		}
		for i, got := range subjects {
			want := &subjectList[i]
			helper.DiffTest(want, got, t)
		}
	})
}
