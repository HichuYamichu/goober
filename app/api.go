package app

import "github.com/hichuyamichu-me/uploader/app/middleware"

func (a *App) setRoutes(mwService *middleware.MiddlewareService) {
	a.router.GET("/files/:name", a.uploadHandler.Download)

	api := a.router.Group("/api")

	filesAPI := api.Group("/files")
	filesAPI.Use(mwService.LoggedIn)
	filesAPI.GET("", a.uploadHandler.FilesInfo)
	filesAPI.POST("", a.uploadHandler.Upload)
	filesAPI.DELETE("/:name", a.uploadHandler.Delete, mwService.Admin)

	userAPI := api.Group("/user")
	userAPI.Use(mwService.LoggedIn)
	userAPI.GET("/list", a.usersHandler.ListUsers, mwService.Admin)
	userAPI.POST("/activate", a.usersHandler.ActivateUser, mwService.Admin)
	userAPI.POST("/password/change", a.usersHandler.ChangePass)
	userAPI.DELETE("/delete/:id", a.usersHandler.DeleteUser)

	authAPI := api.Group("/auth")
	authAPI.POST("/login", a.authHandler.Login)
	authAPI.POST("/register", a.authHandler.Register)

	a.router.Use(mwService.ServeSPA())
}
