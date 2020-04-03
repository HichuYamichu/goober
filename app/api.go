package app

import "github.com/hichuyamichu-me/goober/app/middleware"

func (a *App) setRoutes(mwService *middleware.MiddlewareService) {
	a.router.GET("/files/:name", a.uploadHandler.Download)

	api := a.router.Group("/api")

	filesAPI := api.Group("/files")
	filesAPI.Use(mwService.LoggedIn)
	filesAPI.GET("", a.uploadHandler.FilesInfo)
	filesAPI.POST("", a.uploadHandler.Upload)
	filesAPI.DELETE("/:name", a.uploadHandler.Delete, mwService.Admin)

	userAPI := api.Group("/users")
	userAPI.Use(mwService.LoggedIn)
	userAPI.GET("", a.usersHandler.ListUsers, mwService.Admin)
	userAPI.GET("/activate/:id", a.usersHandler.ActivateUser, mwService.Admin)
	userAPI.GET("/token", a.usersHandler.ChangeToken)
	userAPI.POST("/password/change", a.usersHandler.ChangePass)
	userAPI.DELETE("/:id", a.usersHandler.DeleteUser)

	authAPI := api.Group("/auth")
	authAPI.POST("/login", a.authHandler.Login)
	authAPI.POST("/register", a.authHandler.Register)

	a.router.Use(mwService.ServeSPA())
}
