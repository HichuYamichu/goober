package app

import "github.com/hichuyamichu-me/uploader/app/middleware"

func (a *App) setRoutes() {
	jwtMW := middleware.JWT()

	api := a.router.Group("/api")
	api.GET("/download/:name", a.uploadHandler.Download)

	filesAPI := api.Group("/files")
	filesAPI.Use(jwtMW)
	filesAPI.DELETE("/delete/:name", a.uploadHandler.Delete, middleware.Admin)
	filesAPI.POST("/upload", a.uploadHandler.Upload)
	filesAPI.GET("/list", a.uploadHandler.FilesInfo)

	userAPI := api.Group("/user")
	userAPI.Use(jwtMW)
	userAPI.GET("/list", a.usersHandler.ListUsers, middleware.Admin)
	userAPI.POST("/activate", a.usersHandler.ActivateUser, middleware.Admin)
	userAPI.POST("/password/change", a.usersHandler.ChangePass)
	userAPI.DELETE("/delete/:id", a.usersHandler.DeleteUser)

	authAPI := api.Group("/auth")
	authAPI.POST("/login", a.authHandler.Login)
	authAPI.POST("/register", a.authHandler.Register)

	a.router.Use(middleware.ServeSPA)
}
