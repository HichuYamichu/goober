package app

import (
	"github.com/hichuyamichu-me/uploader/db"
	"github.com/hichuyamichu-me/uploader/router"
	"github.com/hichuyamichu-me/uploader/router/middleware"
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

	uploadService := upload.NewService()
	uploadHandler := upload.NewHandler(uploadService)

	r := router.New()

	api := r.Group("/api")
	api.GET("/download/:name", uploadHandler.Download)
	api.GET("/status", uploadHandler.Status, middleware.JWT)
	api.POST("/login", usersHandler.Login)
	api.POST("/upload", uploadHandler.Upload, middleware.JWT)
	api.POST("/password/change", usersHandler.ChangePass, middleware.JWT)

	adminAPI := api.Group("/admin")
	adminAPI.Use(middleware.JWT)
	adminAPI.Use(middleware.Admin)
	adminAPI.POST("/user", usersHandler.CreateUser)
	adminAPI.PUT("/user", usersHandler.UpdateUser)
	adminAPI.DELETE("/user", usersHandler.DeleteUser)
	adminAPI.DELETE("/delete_file/:name", uploadHandler.Delete, middleware.JWT)

	r.Use(middleware.ServeSPA)

	return r
}
