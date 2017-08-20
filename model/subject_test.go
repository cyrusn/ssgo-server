package model_test

import (
	"testing"

	"github.com/cyrusn/ssgo/model"
)

var subjectList = []model.Subject{
	model.Subject{"bio", 1, "Biology", "生物", 0},
	model.Subject{"bafs", 1, "Business, Accounting and Financial Studies", "企業、會計與財務概論", 0},
	model.Subject{"ict", 2, "Information and Communication Technology", "資訊及通訊科技", 0},
	model.Subject{"econ", 2, "Economics", "經濟", 0},
}

var TestSubjectTable = func(t *testing.T) {
	// t.Run("Create subject table", TestCreateSubjectTable)
	t.Run("Insert Subject List", TestInsertSubject)
	t.Run("Get All Subject", TestAllSubjects)
	t.Run("Get Subject by subject code (bio)", TestGetSubject(0))
	t.Run("Update Subject (bafs)'s Capacity to 20'", TestUpdateSubjectCapacity(1, 20))
	t.Run("Get Subject by subject code (bafs)", TestGetSubject(1))
}

var TestUpdateSubjectCapacity = func(index, capacity int) func(*testing.T) {
	return func(t *testing.T) {
		code := subjectList[index].Code
		if err := db.UpdateSubjectCapacity(code, capacity); err != nil {
			t.Fatal(err)
		}
		subjectList[index].Capacity = capacity
	}
}

var TestInsertSubject = func(t *testing.T) {
	for _, s := range subjectList {
		if err := db.InsertSubject(s); err != nil {
			t.Fatal(err)
		}
	}
}

var TestAllSubjects = func(t *testing.T) {
	subjects, err := db.AllSubjects()
	if err != nil {
		t.Fatal(err)
	}

	for i, got := range subjects {
		want := &subjectList[i]
		diffTest(want, got, t)
	}
}

var TestGetSubject = func(index int) func(*testing.T) {
	return func(t *testing.T) {
		want := subjectList[index]
		got, err := db.GetSubject(want.Code)
		if err != nil {
			t.Fatal(err)
		}
		diffTest(&want, got, t)
	}
}
