package main

import (
	"crypto/rand"
	"fmt"

	"github.com/hichuyamichu-me/uploader/models"
	usersService "github.com/hichuyamichu-me/uploader/services/users"
	"github.com/hichuyamichu-me/uploader/store"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

const bits = 256

func generateKey() string {
	size := bits / 8
	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return string(key)
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	key := generateKey()
	viper.Set("secret_key", string(key))

	store.ConnectDB()
	suName := viper.GetString("super_user_name")
	suPass := viper.GetString("super_user_pass")

	user := usersService.FindOneByUsername(suName)
	if user == nil {
		user = &models.User{Username: suName, Pass: suPass, Admin: true, Quota: int64(100000000)}
		user.Pass, err = usersService.HashPassword(user.Pass)
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
		usersService.Create(user)
		viper.Set("super_user_pass", user.Pass)
	}

	viper.WriteConfig()
}
