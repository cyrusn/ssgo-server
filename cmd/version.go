package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v1.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Subject Selection System Backend Server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfgFile, privateKey, dbPath, isOverwrite, teacherJSONPath, studentJSONPath, subjectJSONPath, port, staticFolderLocation, lifeTime)
		fmt.Println(version)
	},
}
