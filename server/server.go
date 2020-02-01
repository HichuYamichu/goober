package server

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hichuyamichu-me/uploader/handlers"
)

// New creates new server instance
func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	conf := parseConfig()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("uploadDir", conf.UploadDir)
			return next(c)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("uploader", nil)
	p.Use(e)

	api := e.Group("/api")
	api.POST("/upload", handlers.Upload, middleware.JWT([]byte("secret")))
	api.GET("/download/:name", handlers.Download)
	api.GET("/status", handlers.Status)
	api.POST("/login", handlers.Login)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: middleware.DefaultSkipper,
		Root:    "client/public",
		Index:   "index.html",
		HTML5:   true,
	}))
	return e
}
