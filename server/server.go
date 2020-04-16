package server

import (
	"context"
	"fmt"

	"github.com/hichuyamichu-me/goober/db"
	"github.com/hichuyamichu-me/goober/server/middleware"
	"github.com/hichuyamichu-me/goober/upload"
	"github.com/labstack/echo/v4"
)

// Server main server struct
type Server struct {
	router        *echo.Echo
	uploadHandler *upload.Handler
}

// New bootstraps server
func New() *Server {
	db := db.Connect()

	uploadRepo := upload.NewRepository(db)
	uploadService := upload.NewService(uploadRepo)
	uploadHandler := upload.NewHandler(uploadService)

	server := &Server{
		router:        router(),
		uploadHandler: uploadHandler,
	}
	server.setRoutes()
	return server
}

func (a *Server) setRoutes() {
	spa := middleware.ServeSPA()
	jwt := middleware.JWT()
	issuer := middleware.ISS
	basicAuth := middleware.BasicAuth()
	parsePerms := middleware.ParsePermissions
	canRead := middleware.CanRead
	canWrite := middleware.CanWrite
	canDelete := middleware.CanDelete

	a.router.GET("/files/:id", a.uploadHandler.Download)

	api := a.router.Group("/api", jwt, issuer, parsePerms)

	uploadsAPI := api.Group("/uploads")
	uploadsAPI.GET("/:page", a.uploadHandler.Files, canRead)
	uploadsAPI.POST("", a.uploadHandler.Upload, canWrite)
	uploadsAPI.DELETE("", a.uploadHandler.Delete, canDelete)

	frontend := a.router.Group("/")
	frontend.Use(basicAuth)
	frontend.Use(spa)
}

// Shutdown shuts down the server
func (a *Server) Shutdown(ctx context.Context) {
	a.router.Shutdown(ctx)
}

// Start starts the server
func (a *Server) Start(host string, port string) error {
	return a.router.Start(fmt.Sprintf("%s:%s", host, port))
}
