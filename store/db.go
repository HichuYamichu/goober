package store

import (
	"fmt"

	"github.com/hichuyamichu-me/uploader/models"
	usersRepo "github.com/hichuyamichu-me/uploader/store/users"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func ConnectDB() {
	dbHost := viper.GetString("db_host")
	dbPort := viper.GetString("db_port")
	dbUser := viper.GetString("db_user")
	dbName := viper.GetString("db_name")
	dbPass := viper.GetString("db_pass")
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&models.User{})
	usersRepo.InjectDB(db)
}
