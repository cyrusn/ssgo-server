package model

import "testing"

var studentRanks = []Rank{
	Rank{"student1", 3},
	Rank{"student2", 1},
	Rank{"student3", 2},
}

var TestRankTable = func(t *testing.T) {
	t.Run("Create rank table", TestCreateRankTable)
	t.Run("Insert rank data", TestInsertRankTable)
	t.Run("Get student2 rank", TestGetStudentRank(1))
	t.Run("Truncate rank table", TestTruncateRankTable)
	t.Run("Insert rank data", TestInsertRankTable)
	t.Run("Get student1 rank", TestGetStudentRank(0))
}

var TestTruncateRankTable = func(t *testing.T) {
	if err := db.TruncateRankTable(); err != nil {
		t.Fatal(err)
	}
}

var TestGetStudentRank = func(index int) func(*testing.T) {
	return func(t *testing.T) {
		rank, err := db.GetStudentRank(studentRanks[index].Username)
		if err != nil {
			t.Fatal(err)
		}

		want := &studentRanks[index]
		got := rank

		diffTest(want, got, t)
	}
}

var TestInsertRankTable = func(t *testing.T) {
	for _, rank := range studentRanks {
		if err := db.InsertStudentRank(rank); err != nil {
			t.Fatal(err)
		}
	}
}

var TestCreateRankTable = func(t *testing.T) {
	if err := db.CreateRankTable(); err != nil {
		t.Fatal(err)
	}
}
