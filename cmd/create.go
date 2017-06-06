package cmd

import (
	"fmt"
	"os"

	"github.com/cyrusn/ssgo/model"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create database for Subject Selection System Backend Server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := model.CreateDBFile(dbPath, isOverwrite); err != nil {
			fmt.Println(err)
			fmt.Println("Please use \"-o\" flag to overwrite existing database")
			os.Exit(1)
		}
		fmt.Println("Database created")
	},
}
