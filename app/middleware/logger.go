package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger logging middleware
func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339_nano}][${method}][${uri}][${status}][${error}]` + "\n",
		Output: os.Stdout,
	})
}
