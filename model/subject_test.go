package model_test

import (
	"fmt"
	"testing"

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
			if err := s.Insert(); err != nil {
				t.Fatal(err)
			}
		})
	}

	t.Run("Update Capacity", func(t *testing.T) {
		for i, subject := range subjectList {
			capacity := 20
			if err := subject.UpdateCapacity(capacity); err != nil {
				t.Fatal(err)
			}
			subjectList[i].Capacity = capacity
		}
	})

	for i, subject := range subjectList {
		name := fmt.Sprintf("Get Each Subjects #%d", i+1)
		t.Run(name, func(t *testing.T) {
			want := subject
			s := new(model.Subject)
			s.Code = want.Code
			if err := s.Get(); err != nil {
				t.Fatal(err)
			}
			diffTest(&want, s, t)
		})
	}

	t.Run("Get All Subjects", func(t *testing.T) {
		subjects, err := model.AllSubjects()
		if err != nil {
			t.Fatal(err)
		}

		for i, got := range subjects {
			want := &subjectList[i]
			diffTest(want, got, t)
		}
	})

}
