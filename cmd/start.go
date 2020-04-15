package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hichuyamichu-me/goober/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts goober's http server",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.New()

		go func() {
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-done
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			server.Shutdown(ctx)
		}()

		host := viper.GetString("goober.host")
		port := viper.GetString("goober.port")
		log.Fatal(server.Start(host, port))
	},
}
