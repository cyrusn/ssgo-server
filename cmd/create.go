package cmd

import (
	"fmt"
	"os"

	"ssgo-server/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create database for Subject Selection System Backend Server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := model.CreateDatabase(dsn, isOverwrite); err != nil {
			fmt.Println("Please use \"-o\" flag to overwrite existing database")
			os.Exit(1)
		}

		fmt.Println("Database created")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().BoolVarP(
		&isOverwrite,
		"overwrite",
		"o",
		DEFAULT_OVERWRITE,
		"overwrite database if database location exist",
	)

	viper.BindPFlags(createCmd.PersistentFlags())
}
