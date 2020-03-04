package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func Connect() *gorm.DB {
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
