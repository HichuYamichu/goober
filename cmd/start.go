package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hichuyamichu-me/uploader/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts uploader's http server",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.New()

		go func() {
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-done
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			app.Shutdown(ctx)
		}()

		port := viper.GetString("port")
		host := viper.GetString("host")
		app.Logger.Fatal(app.Start(fmt.Sprintf("%s:%s", host, port)))
	},
}
