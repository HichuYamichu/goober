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

	filesAPI := api.Group("/files")
	filesAPI.Use(jwtMW)
	filesAPI.DELETE("/delete/:name", uploadHandler.Delete, middleware.Admin)
	filesAPI.POST("/upload", uploadHandler.Upload)
	filesAPI.GET("/list", uploadHandler.FilesInfo)

	userAPI := api.Group("/user")
	userAPI.Use(jwtMW)
	userAPI.GET("/list", usersHandler.ListUsers, middleware.Admin)
	userAPI.POST("/activate", usersHandler.ActivateUser, middleware.Admin)
	userAPI.POST("/password/change", usersHandler.ChangePass)
	userAPI.DELETE("/delete/:id", usersHandler.DeleteUser)

	authAPI := api.Group("/auth")
	authAPI.POST("/login", authHandler.Login)
	authAPI.POST("/register", authHandler.Register)

	r.Use(middleware.ServeSPA)

	return r
}
