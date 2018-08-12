package cmd

import (
	"fmt"
	"os"

	"github.com/cyrusn/ssgo-server/model/subject"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

func init() {
	importCmd.AddCommand(subjectCmd)

	subjectCmd.PersistentFlags().StringVarP(
		&subjectJSONPath,
		"subject",
		"s",
		SUBJECT_JSON_PATH,
		"path of subject.json file\nplease check README.md for the schema",
	)

	viper.BindPFlags(subjectCmd.PersistentFlags())
}
