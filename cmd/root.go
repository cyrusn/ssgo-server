package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ssgo",
	Short: "Welcome to Subject Selection System Backend Server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"c",
		CONFIG_PATH,
		"config file",
	)

	rootCmd.PersistentFlags().StringVarP(
		&privateKey,
		"key",
		"k",
		PRIVATE_KEY,
		"change the private key for authentication on jwt",
	)

	rootCmd.PersistentFlags().StringVarP(
		&DSN,
		"dsn",
		"d",
		DEFAULT_DSN,
		"Data source name of mysql. [ref https://github.com/go-sql-driver/mysql]",
	)

	for _, name := range []string{"key", "database"} {
		viper.BindPFlag(name, rootCmd.PersistentFlags().Lookup(name))
	}
}
