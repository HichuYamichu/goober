package cmd

import (
	"github.com/hichuyamichu-me/goober/db"
	"github.com/hichuyamichu-me/goober/upload"
	"github.com/spf13/cobra"
)

const bits = 256

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs migrations and ensures admin user existance",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Connect()
		defer db.Close()
		db.DropTableIfExists(&upload.File{})
		db.AutoMigrate(&upload.File{})
	},
}
