package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "goober",
		Short: "Simple file server",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/goober/")
	viper.AddConfigPath("$HOME/.goober")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "3000")
	viper.SetDefault("domain", fmt.Sprintf("%s:%s", viper.Get("host"), viper.Get("port")))
	viper.SetDefault("upload_dir", "./data")
	viper.SetDefault("frontend", true)

	err = verifyConfig()
	if err != nil {
		log.Fatalf("error validating config: %s", err)
	}

	rootCmd.AddCommand(startCmd)
}

func verifyConfig() error {
	if viper.IsSet("admin") {
		users := viper.GetStringSlice("admin")
		for _, user := range users {
			split := strings.Split(user, ":")
			if len(split) != 2 {
				return fmt.Errorf("admin value must be of format username:password")
			}
		}
	}

	if viper.IsSet("jwt") {
		if !viper.IsSet("jwt.type") {
			return fmt.Errorf("jwt.type must be set when jwt is enabled")
		}
		if !viper.IsSet("jwt.key") {
			return fmt.Errorf("jwt.key must be set when jwt is enabled")
		}
	}

	return nil
}
