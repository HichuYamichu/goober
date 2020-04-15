package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func Connect() *gorm.DB {
	dbHost := viper.GetString("postgres.host")
	dbPort := viper.GetString("postgres.port")
	dbUser := viper.GetString("postgres.user")
	dbName := viper.GetString("postgres.name")
	dbPass := viper.GetString("postgres.pass")
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
