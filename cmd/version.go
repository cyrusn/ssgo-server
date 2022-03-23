package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v1.7.0"
const releaseHistory = `Release History:
	v1.7.0: remove serving statics site and use docker on deployment
	v1.6.0: parent signature is saved to database.
	v1.5.0: record students who are going to take HMSC as 3rd elective.
	v1.4.2: better format
	v1.3.0: add timestamp feature
	v1.2.0: use mysql instead of sqlite3
	v1.0.0: first release
`

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Subject Selection System Backend Server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("current version:", version)
		fmt.Println(releaseHistory)
	},
}
