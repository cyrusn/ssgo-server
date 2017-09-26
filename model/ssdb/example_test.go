package ssdb_test

import "log"

func ExampleInitDB() {
	path := "./testing.db"
	db, err := model.InitDB(path)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
