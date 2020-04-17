package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// ServeSPA middleware for serving spa
func ServeSPA() echo.MiddlewareFunc {
	skipper := func(echo.Context) bool { return !viper.GetBool("frontend") }
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: skipper,
		Root:    "web/public/",
		Index:   "index.html",
		HTML5:   true,
		Browse:  false,
	})
}
