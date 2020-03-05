package app

import (
	"github.com/hichuyamichu-me/uploader/app/middleware"
	"github.com/hichuyamichu-me/uploader/auth"
	"github.com/hichuyamichu-me/uploader/db"
	"github.com/hichuyamichu-me/uploader/upload"
	"github.com/hichuyamichu-me/uploader/users"
	"github.com/labstack/echo/v4"
)

// New bootstraps app
func New() *echo.Echo {
	db := db.Connect()

	usersRepo := users.NewRepository(db)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	authService := auth.NewService(usersRepo)
	authHandler := auth.NewHandler(authService, usersService)

	uploadService := upload.NewService()
	uploadHandler := upload.NewHandler(uploadService)

	r := newRouter()

	jwtMW := middleware.JWT()

	api := r.Group("/api")
	api.GET("/download/:name", uploadHandler.Download)
	api.GET("/status", uploadHandler.Status, jwtMW)
	api.POST("/upload", uploadHandler.Upload, jwtMW)

	usersAPI := api.Group("/user")
	usersAPI.Use(jwtMW)
	usersAPI.POST("/password/change", usersHandler.ChangePass)

	authAPI := api.Group("/auth")
	authAPI.POST("/login", authHandler.Login)
	authAPI.POST("/register", authHandler.Register)

	adminAPI := api.Group("/admin")
	adminAPI.Use(jwtMW)
	adminAPI.Use(middleware.Admin)
	adminAPI.PUT("/activate", usersHandler.ActivateUser)
	adminAPI.DELETE("/delete_file/:name", uploadHandler.Delete)

	r.Use(middleware.ServeSPA)

	return r
}
