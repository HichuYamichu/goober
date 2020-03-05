package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Recover echo Recover re-export
func Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}
