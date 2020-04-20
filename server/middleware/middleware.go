package middleware

import (
	"strings"

	"github.com/spf13/viper"
)

type skipper struct{}

func pathSkip(path string) bool {
	return serveFilesSkip(path) || spaSkip(path)
}

func serveFilesSkip(path string) bool {
	return viper.GetBool("skip_serving_auth") && path == "/files/:id"
}

func spaSkip(path string) bool {
	return viper.GetBool("skip_frontend_auth") && path != "/files/:id" && !strings.HasPrefix(path, "/api")
}
