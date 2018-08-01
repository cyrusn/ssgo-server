package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Subject Selection System Backend Server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
