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

	applyDefaultConfig()
	err = verifyConfig()
	if err != nil {
		log.Fatalf("error validating config: %s", err)
	}

	rootCmd.AddCommand(startCmd)
}

func applyDefaultConfig() {
	if !viper.IsSet("host") {
		viper.Set("host", "localhost")
	}

	if !viper.IsSet("port") {
		viper.Set("port", "3000")
	}

	if !viper.IsSet("domain") {
		viper.Set("domain", fmt.Sprintf("%s:%s", viper.Get("host"), viper.Get("port")))
	}

	if !viper.IsSet("upload_dir") {
		viper.Set("upload_dir", "./data")
	}

	if !viper.IsSet("frontend") {
		viper.Set("frontend", true)
	}
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
