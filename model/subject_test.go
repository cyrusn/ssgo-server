package model_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var subjectList = []subject.Subject{
	model.Subject{"bio", 1, "Biology", "生物", 0},
	model.Subject{"bafs", 1, "Business, Accounting and Financial Studies", "企業、會計與財務概論", 0},
	model.Subject{"ict", 2, "Information and Communication Technology", "資訊及通訊科技", 0},
	model.Subject{"econ", 2, "Economics", "經濟", 0},
}

var TestUpdateCapacity = func(index, capacity int) func(*testing.T) {
	return func(t *testing.T) {
		code := subjectList[index].Code
		if err := subject.UpdateCapacity(db, code, capacity); err != nil {
			t.Fatal(err)
		}
		subjectList[index].Capacity = capacity
	}
}

var TestInsert = func(t *testing.T) {
	for _, s := range subjectList {
		if err := subject.Insert(db, s); err != nil {
			t.Fatal(err)
		}
	}
}

var TestAll = func(t *testing.T) {
	subjects, err := subject.All(db)
	if err != nil {
		t.Fatal(err)
	}

	for i, got := range subjects {
		want := &subjectList[i]
		helper.DiffTest(want, got, t)
	}
}

var TestGet = func(index int) func(*testing.T) {
	return func(t *testing.T) {
		want := subjectList[index]
		got, err := subject.Get(db, want.Code)
		if err != nil {
			t.Fatal(err)
		}
		helper.DiffTest(&want, got, t)
	}
}
