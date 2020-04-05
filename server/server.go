package server

import (
	"context"
	"fmt"

	"github.com/hichuyamichu-me/goober/db"
	"github.com/hichuyamichu-me/goober/domain/auth"
	"github.com/hichuyamichu-me/goober/domain/upload"
	"github.com/hichuyamichu-me/goober/domain/users"
	"github.com/hichuyamichu-me/goober/server/middleware"
	"github.com/labstack/echo/v4"
)

// Server main server struct
type Server struct {
	router *echo.Echo

	usersHandler  *users.Handler
	authHandler   *auth.Handler
	uploadHandler *upload.Handler

	middlewareService *middleware.Service
}

// New bootstraps server
func New() *Server {
	db := db.Connect()

	usersRepo := users.NewRepository(db)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	authService := auth.NewService(usersRepo)
	authHandler := auth.NewHandler(authService, usersService)

	uploadRepo := upload.NewRepository(db)
	uploadService := upload.NewService(uploadRepo)
	uploadHandler := upload.NewHandler(uploadService)

	mwService := middleware.NewService(usersRepo)

	server := &Server{
		router:            newRouter(),
		usersHandler:      usersHandler,
		authHandler:       authHandler,
		uploadHandler:     uploadHandler,
		middlewareService: mwService,
	}

	server.setRoutes()

	return server
}

// Shutdown shuts down the server
func (a *Server) Shutdown(ctx context.Context) {
	a.router.Shutdown(ctx)
}

// Start starts the server
func (a *Server) Start(host string, port string) error {
	return a.router.Start(fmt.Sprintf("%s:%s", host, port))
}
