package server

import (
	"context"
	"fmt"

	"github.com/hichuyamichu-me/goober/db"
	"github.com/hichuyamichu-me/goober/server/middleware"
	"github.com/hichuyamichu-me/goober/upload"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Server main server struct
type Server struct {
	router        *echo.Echo
	db            *gorm.DB
	uploadHandler *upload.Handler
}

// New bootstraps server
func New() (*Server, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	uploadRepo := upload.NewRepository(db)
	uploadService := upload.NewService(uploadRepo)
	uploadHandler := upload.NewHandler(uploadService)

	server := &Server{
		router:        echo.New(),
		db:            db,
		uploadHandler: uploadHandler,
	}
	server.configure()
	server.setRoutes()
	return server, nil
}

func (s *Server) configure() {
	s.router.HideBanner = true
	s.router.HTTPErrorHandler = httpErrorHandler
	s.router.Validator = NewValidator()
	s.router.Logger.SetLevel(log.INFO)

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.BodyLimit())
}

func (s *Server) setRoutes() {
	spa := middleware.ServeSPA()
	jwt := middleware.JWT
	issuer := middleware.ISS
	basicAuth := middleware.BasicAuth()
	canRead := middleware.CanRead
	canWrite := middleware.CanWrite
	canDelete := middleware.CanDelete

	s.router.Use(jwt, issuer, basicAuth)
	s.router.GET("/files/:id", s.uploadHandler.Download, canRead)

	api := s.router.Group("/api")
	uploadsAPI := api.Group("/uploads")
	uploadsAPI.GET("/:page", s.uploadHandler.Files, canRead)
	uploadsAPI.POST("", s.uploadHandler.Upload, canWrite)
	uploadsAPI.DELETE("", s.uploadHandler.Delete, canDelete)

	s.router.Use(spa)
}

// Shutdown shuts down the server
func (s *Server) Shutdown(ctx context.Context) {
	s.router.Shutdown(ctx)
	s.db.Close()
}

// Start starts the server
func (s *Server) Start(host string, port string) error {
	return s.router.Start(fmt.Sprintf("%s:%s", host, port))
}
