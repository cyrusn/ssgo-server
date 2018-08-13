package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func checkPathExist(paths ...string) {
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatalf(`"%s" doesn't exist`, path)
		}
	}
}

func unmarshalJSON(jsonFilePath string, v interface{}) {
	content := readJSONFile(jsonFilePath)
	if err := json.Unmarshal(content, &v); err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
}

func readJSONFile(jsonFilePath string) []byte {
	file, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	return file
}

func openDB(DSN string) *sql.DB {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
