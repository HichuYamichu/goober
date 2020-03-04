package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "uploader",
		Short: "Private upload server",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(migrateCmd)
}
