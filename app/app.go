package app

import (
	"context"
	"fmt"

	"github.com/hichuyamichu-me/goober/app/middleware"
	"github.com/hichuyamichu-me/goober/db"
	"github.com/hichuyamichu-me/goober/internal/auth"
	"github.com/hichuyamichu-me/goober/internal/upload"
	"github.com/hichuyamichu-me/goober/internal/users"
	"github.com/labstack/echo/v4"
)

// App main app struct
type App struct {
	usersHandler  *users.Handler
	authHandler   *auth.Handler
	uploadHandler *upload.Handler

	router *echo.Echo
}

// New bootstraps app
func New() *App {
	db := db.Connect()

	usersRepo := users.NewRepository(db)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	authService := auth.NewService(usersRepo)
	authHandler := auth.NewHandler(authService, usersService)

	uploadService := upload.NewService()
	uploadHandler := upload.NewHandler(uploadService)

	app := &App{
		router:        newRouter(),
		usersHandler:  usersHandler,
		authHandler:   authHandler,
		uploadHandler: uploadHandler,
	}

	mwService := middleware.NewMiddlewareService(usersRepo)
	app.setRoutes(mwService)

	return app
}

// Shutdown shuts down the app
func (a *App) Shutdown(ctx context.Context) {
	a.router.Shutdown(ctx)
}

// Start starts the app
func (a *App) Start(host string, port string) error {
	return a.router.Start(fmt.Sprintf("%s:%s", host, port))
}
