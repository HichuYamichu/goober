package server

import (
	"text/template"

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
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("uploader", nil)
	p.Use(e)

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Renderer = t
	e.POST("/upload", handlers.Upload)
	e.GET("/download/:name", handlers.Download)
	e.GET("/status", handlers.Status)
	e.GET("/login", handlers.ServeLogin)
	e.POST("/login", handlers.Login)
	e.GET("/", handlers.Index)

	return e
}
