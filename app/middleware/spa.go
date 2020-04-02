package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ServeSPA middleware for serving spa
func (mws *MiddlewareService) ServeSPA() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: middleware.DefaultSkipper,
		Root:    "web/public/",
		Index:   "index.html",
		HTML5:   true,
		Browse:  false,
	})
}
