package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func Logger() echo.MiddlewareFunc {
	return middleware.Logger()
}

func BodyLimit() echo.MiddlewareFunc {
	return middleware.BodyLimit(viper.GetString("max_body_size"))
}

func Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

func pathSkip(path string) bool {
	return serveFilesSkip(path) || spaSkip(path)
}

func serveFilesSkip(path string) bool {
	return viper.GetBool("skip_serving_auth") && path == "/files/:id"
}

func spaSkip(path string) bool {
	return viper.GetBool("skip_frontend_auth") && path != "/files/:id" && !strings.HasPrefix(path, "/api")
}
