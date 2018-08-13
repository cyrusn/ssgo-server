package cmd

import (
	"fmt"
	"os"

	"github.com/cyrusn/ssgo-server/model/auth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var teacherCmd = &cobra.Command{
	Use:   "teacher",
	Short: "Import teachers to Credential table in database",
	Run: func(cmd *cobra.Command, args []string) {
		var credentials []auth.Credential

		checkPathExist(teacherJSONPath)
		unmarshalJSON(teacherJSONPath, &credentials)

		db := &auth.DB{openDB(DSN), &secret}
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

func init() {
	importCmd.AddCommand(teacherCmd)

	teacherCmd.PersistentFlags().StringVarP(
		&teacherJSONPath,
		"teacher",
		"t",
		TEACHER_JSON_PATH,
		"path of teacher.json file\nplease check README.md for the schema",
	)

	viper.BindPFlags(teacherCmd.PersistentFlags())
}
