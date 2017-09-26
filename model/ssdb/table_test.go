package ssdb_test

import "testing"

var PanicTestCreateTables = func(t *testing.T) {
	expectError("CreateTables", t, func() {
		if err := db.CreateTables(); err != nil {
			panic(err)
		}
	})
}
var TestCreateTable = func(t *testing.T) {
	err := db.CreateTables()
	if err != nil {
		t.Fatal(err)
	}
}
