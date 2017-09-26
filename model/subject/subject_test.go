package subject_test

import (
	"log"
	"testing"

	"github.com/cyrusn/ssgo/model/helper"
	"github.com/cyrusn/ssgo/model/ssdb"
	"github.com/cyrusn/ssgo/model/subject"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DBPath = "../../database/test.db"
)

var db *ssdb.DB

func init() {
	ssdb.Cleanup(DBPath)
	var err error

	db, err = ssdb.InitDB(DBPath)
	if err != nil {
		log.Fatal(err)
	}
	db.CreateTables()
}

var subjectList = []subject.Subject{
	subject.Subject{"bio", 1, "Biology", "生物", 0},
	subject.Subject{"bafs", 1, "Business, Accounting and Financial Studies", "企業、會計與財務概論", 0},
	subject.Subject{"ict", 2, "Information and Communication Technology", "資訊及通訊科技", 0},
	subject.Subject{"econ", 2, "Economics", "經濟", 0},
}

func Test(t *testing.T) {
	// t.Run("Create subject table", TestCreateSubjectTable)
	t.Run("Insert subjects", TestInsert)
	t.Run("Get all subjects", TestAll)
	t.Run("Get subject by subject code (bio)", TestGet(0))
	t.Run("Update bafs's Capacity to 20'", TestUpdateCapacity(1, 20))
	t.Run("check if bafs's capacity updated", TestGet(1))
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
