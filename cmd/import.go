package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cyrusn/ssgo/model/auth"
	"github.com/cyrusn/ssgo/model/student"
	"github.com/cyrusn/ssgo/model/subject"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import users and subjects to database",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var teacherCmd = &cobra.Command{
	Use:   "teacher",
	Short: "Import teachers to Credential table in database",
	Run: func(cmd *cobra.Command, args []string) {
		var credentials []auth.Credential

		checkPathExist(dbPath, teacherJSONPath)
		unmarshalJSON(teacherJSONPath, &credentials)

		db := &auth.DB{openDB(dbPath), &secret}
		defer db.Close()

		for _, c := range credentials {
			if err := db.Insert(&c); err != nil {
				fmt.Printf("Import error: %v\n", err)
				os.Exit(1)
			}
		}
		fmt.Println("teachers are imported")
	},
}

var subjectCmd = &cobra.Command{
	Use:   "subject",
	Short: "Import subjects to Subject table in database",
	Run: func(cmd *cobra.Command, args []string) {
		var codes []string

		checkPathExist(dbPath, subjectJSONPath)
		unmarshalJSON(subjectJSONPath, &codes)

		db := &subject.DB{openDB(dbPath)}
		defer db.Close()

		for _, c := range codes {
			s := subject.Subject{Code: c}
			if err := db.Insert(&s); err != nil {
				fmt.Printf("Import error: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Println("subjects are imported")
	},
}

var studentCmd = &cobra.Command{
	Use:   "student",
	Short: "Import students to Credential and Student table in database",
	Run: func(cmd *cobra.Command, args []string) {
		var students []student.Student
		var credentials []auth.Credential
		checkPathExist(dbPath, studentJSONPath)
		unmarshalJSON(studentJSONPath, &students)
		unmarshalJSON(studentJSONPath, &credentials)

		db := openDB(dbPath)
		defer db.Close()

		credentialDB := &auth.DB{db, &secret}
		studentDB := &student.DB{db}

		for _, c := range credentials {
			c.Role = "STUDENT"

			if err := credentialDB.Insert(&c); err != nil {
				fmt.Printf("Import error: %v\n", err)
				os.Exit(1)
			}
		}

		for _, s := range students {
			s.Priority = []int{}
			s.IsConfirmed = false
			s.Rank = -1

			if err := studentDB.Insert(&s); err != nil {
				fmt.Printf("Import error: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Println("students are imported")
	},
}

func openDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	return db
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
