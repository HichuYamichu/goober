package cmd

import (
	"crypto/rand"
	"log"

	"github.com/hichuyamichu-me/uploader/db"
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
			log.Fatal(err)
		}
		viper.Set("secret_key", string(key))
		viper.WriteConfig()

		db := db.Connect()
		db.DropTableIfExists(&users.User{})
		db.AutoMigrate(&users.User{})

		usersRepo := users.NewRepository(db)
		usersService := users.NewService(usersRepo)

		user := &users.User{Username: "root", Pass: "root", Admin: true, Active: true, Quota: int64(100000000)}
		pass, err := usersService.HashPassword(user.Pass)
		if err != nil {
			log.Fatal(err)
		}
		user.Pass = pass
		err = usersRepo.Create(user)
		if err != nil {
			log.Fatal(err)
		}
	},
}
