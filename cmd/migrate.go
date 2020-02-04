package cmd

import (
	"crypto/rand"
	"log"

	"github.com/hichuyamichu-me/uploader/internal/users"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const bits = 256

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs migrations, generates secure key and ensures admin user existance",
	Run: func(cmd *cobra.Command, args []string) {
		size := bits / 8
		key := make([]byte, size)
		if _, err := rand.Read(key); err != nil {
			panic(err)
		}
		viper.Set("secret_key", string(key))
		viper.WriteConfig()

		db := connectDB()
		db.DropTableIfExists(&users.User{})
		db.AutoMigrate(&users.User{})

		cache := connectCache()
		usersRepo := users.NewRepository(db)
		usersService := users.NewService(usersRepo, cache)

		suName := viper.GetString("super_user_name")
		suPass := viper.GetString("super_user_pass")

		user := &users.User{Username: suName, Pass: suPass, Admin: true, Quota: int64(100000000)}
		err := usersService.CreateUser(user)
		if err != nil {
			log.Fatal(err)
		}
	},
}
