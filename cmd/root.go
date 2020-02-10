package cmd

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func connectDB() *gorm.DB {
	dbHost := viper.GetString("db_host")
	dbPort := viper.GetString("db_port")
	dbUser := viper.GetString("db_user")
	dbName := viper.GetString("db_name")
	dbPass := viper.GetString("db_pass")
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	return db
}
