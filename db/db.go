package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

func Connect() *gorm.DB {
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbUser := viper.GetString("db.user")
	dbName := viper.GetString("db.name")
	dbPass := viper.GetString("db.pass")

	dbType := strings.ToLower(viper.GetString("db.type"))
	var connStr string
	switch dbType {
	case "postgres":
		connStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	case "mysql":
		connStr = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	case "mssql":
		connStr = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)
	default:
		dbType = "sqlite3"
		connStr = "./goober.db"
	}
	db, err := gorm.Open(dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
