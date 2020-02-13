package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hichuyamichu-me/uploader/internal/upload"
	"github.com/hichuyamichu-me/uploader/internal/users"
	"github.com/spf13/viper"
)

// New creates new server instance
func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("uploader", nil)
	p.Use(e)

	jwtMiddleware := middleware.JWT([]byte(viper.GetString("secret_key")))

	usersRepo := users.NewRepository(db)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	uploadService := upload.NewService()
	uploadHandler := upload.NewHandler(uploadService)

	api := e.Group("/api")
	api.GET("/download/:name", uploadHandler.Download)
	api.GET("/status", uploadHandler.Status, jwtMiddleware)
	api.POST("/login", usersHandler.Login)
	api.POST("/upload", uploadHandler.Upload, jwtMiddleware)
	api.POST("/password/change", usersHandler.ChangePass, jwtMiddleware)

	adminAPI := api.Group("/admin")
	adminAPI.Use(jwtMiddleware)
	adminAPI.Use(adminMiddleware)
	adminAPI.POST("/user", usersHandler.CreateUser)
	adminAPI.PUT("/user", usersHandler.UpdateUser)
	adminAPI.DELETE("/user", usersHandler.DeleteUser)
	adminAPI.DELETE("/delete_file/:name", uploadHandler.Delete, jwtMiddleware)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: middleware.DefaultSkipper,
		Root:    "web/public/",
		Index:   "index.html",
		HTML5:   true,
		Browse:  false,
	}))
	return e
}

func adminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)
		if !isAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
