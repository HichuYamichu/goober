package cmd

import (
	"log"

	"github.com/hichuyamichu-me/goober/db"
	"github.com/hichuyamichu-me/goober/internal/users"
	"github.com/spf13/cobra"
)

const bits = 256

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs migrations, generates secure key and ensures admin user existance",
	Run: func(cmd *cobra.Command, args []string) {
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

		token, err := usersService.GenerateToken(user.Username)
		if err != nil {
			log.Fatal(err)
		}
		user.Token = token

		err = usersRepo.Create(user)
		if err != nil {
			log.Fatal(err)
		}
	},
}
