package cmd

import (
	"fmt"
	"os"

	"github.com/cyrusn/ssgo-server/model/auth"
	"github.com/cyrusn/ssgo-server/model/student"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var studentCmd = &cobra.Command{
	Use:   "student",
	Short: "Import students to Credential and Student table in database",
	Run: func(cmd *cobra.Command, args []string) {
		var students []student.Student
		var credentials []auth.Credential
		checkPathExist(studentJSONPath)
		unmarshalJSON(studentJSONPath, &students)
		unmarshalJSON(studentJSONPath, &credentials)

		db := openDB(DSN)
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
			s.Priorities = []int{}
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

func init() {
	importCmd.AddCommand(studentCmd)

	studentCmd.PersistentFlags().StringVarP(
		&studentJSONPath,
		"student",
		"u",
		STUDENT_JSON_PATH,
		"path of student.json file\nplease check README.md for the schema",
	)

	viper.BindPFlags(studentCmd.PersistentFlags())
}
