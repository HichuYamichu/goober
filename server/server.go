package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	uploadHandler "github.com/hichuyamichu-me/uploader/handlers/upload"
	userHandler "github.com/hichuyamichu-me/uploader/handlers/users"
	"github.com/spf13/viper"
)

func ensureClaim(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			isAdmin := claims[key].(bool)
			if !isAdmin {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}

// New creates new server instance
func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("uploader", nil)
	p.Use(e)

	api := e.Group("/api")
	api.POST("/upload", uploadHandler.Upload, middleware.JWT([]byte(viper.GetString("secret_key"))))
	api.GET("/download/:name", uploadHandler.Download)
	api.GET("/status", uploadHandler.Status, middleware.JWT([]byte(viper.GetString("secret_key"))), ensureClaim("read"))
	api.POST("/login", userHandler.Login)
	api.POST("/register/:inviteID", userHandler.Register)

	adminAPI := api.Group("/admin")
	adminAPI.Use(middleware.JWT([]byte(viper.GetString("secret_key"))))
	adminAPI.Use(ensureClaim("admin"))
	adminAPI.POST("/invite", userHandler.Invite)
	adminAPI.DELETE("/user", userHandler.DeleteUser)
	adminAPI.PUT("/user", userHandler.UpdateUser)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: middleware.DefaultSkipper,
		Root:    "client/public",
		Index:   "index.html",
		HTML5:   true,
	}))
	return e
}
