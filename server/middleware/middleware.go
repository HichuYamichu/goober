package middleware

import (
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
