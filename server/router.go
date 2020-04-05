package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// newRouter creates new preconfigured router
func newRouter() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = httpErrorHandler
	e.Validator = NewValidator()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
