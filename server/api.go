package server

func (a *Server) setRoutes() {
	a.router.GET("/files/:id", a.uploadHandler.Download)

	api := a.router.Group("/api")

	filesAPI := api.Group("/uploads")
	filesAPI.Use(a.middlewareService.LoggedIn)
	filesAPI.GET("/:page", a.uploadHandler.Files)
	filesAPI.POST("", a.uploadHandler.Upload)
	filesAPI.DELETE("", a.uploadHandler.Delete, a.middlewareService.Admin)

	userAPI := api.Group("/users")
	userAPI.Use(a.middlewareService.LoggedIn)
	userAPI.GET("", a.usersHandler.ListUsers, a.middlewareService.Admin)
	userAPI.GET("/activate/:id", a.usersHandler.ActivateUser, a.middlewareService.Admin)
	userAPI.GET("/token", a.usersHandler.ChangeToken)
	userAPI.POST("/password/change", a.usersHandler.ChangePass)
	userAPI.DELETE("/:id", a.usersHandler.DeleteUser)

	authAPI := api.Group("/auth")
	authAPI.POST("/login", a.authHandler.Login)
	authAPI.POST("/register", a.authHandler.Register)

	a.router.Use(a.middlewareService.ServeSPA())
}
