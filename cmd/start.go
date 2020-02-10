package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hichuyamichu-me/uploader/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts uploader's http server",
	Run: func(cmd *cobra.Command, args []string) {
		db := connectDB()
		srv := server.New(db)

		go func() {
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-done
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
		}()

		port := viper.GetString("port")
		host := viper.GetString("host")
		srv.Logger.Fatal(srv.Start(fmt.Sprintf("%s:%s", host, port)))
	},
}
