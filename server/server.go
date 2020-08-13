package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hichuyamichu-me/goober/errors"
	"github.com/hichuyamichu-me/goober/files"
	"github.com/hichuyamichu-me/goober/server/middleware"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

// Server main server struct
type Server struct {
	router     *echo.Echo
	uploadRepo *files.Repository
	db         *gorm.DB
	domain     string
}

// New bootstraps server
func New(db *gorm.DB) (*Server, error) {
	s := &Server{
		router:     echo.New(),
		uploadRepo: files.NewRepository(db),
		db:         db,
		domain:     viper.GetString("domain"),
	}

	s.router.HideBanner = true
	s.router.HTTPErrorHandler = httpErrorHandler
	s.router.Validator = NewValidator()
	s.router.Logger.SetLevel(log.INFO)

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.BodyLimit())

	jwt := middleware.JWT
	issuer := middleware.ISS
	basicAuth := middleware.BasicAuth()
	canRead := middleware.CanRead
	canWrite := middleware.CanWrite
	canDelete := middleware.CanDelete

	s.router.Use(jwt, issuer, basicAuth)
	s.router.GET("/:id", s.Download, canRead)

	s.router.GET("/list/:page", s.Files, canRead)
	s.router.POST("", s.Upload, canWrite)
	s.router.DELETE("", s.Delete, canDelete)

	return s, nil
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

func (s *Server) Upload(c echo.Context) error {
	const op errors.Op = "upload/handler.Upload"

	form, err := c.MultipartForm()
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}
	files := form.File["files"]

	type upload struct {
		URL    string `json:"url,omitempty"`
		Name   string `json:"name"`
		Size   int64  `json:"size"`
		Reason error  `json:"reason,omitempty"`
	}

	type uploadResult struct {
		Succeeded []*upload `json:"succeeded"`
		Failed    []*upload `json:"failed"`
	}

	succeeded := make([]*upload, 0)
	failed := make([]*upload, 0)
	for _, file := range files {
		id, err := s.uploadRepo.Save(file)
		if err != nil {
			u := &upload{Name: file.Filename, Size: file.Size, Reason: err}
			failed = append(failed, u)
		}
		u := &upload{
			URL:  fmt.Sprintf("https://%s/%s", s.domain, id),
			Name: file.Filename,
			Size: file.Size,
		}
		succeeded = append(succeeded, u)
	}

	res := &uploadResult{Succeeded: succeeded, Failed: failed}
	return c.JSON(http.StatusOK, res)
}

func (s *Server) Download(c echo.Context) error {
	const op errors.Op = "upload/handler.Download"

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	f, err := s.uploadRepo.Get(id)
	if err != nil {
		return errors.E(err, op)
	}
	defer f.Inner.Close()

	headerVal := fmt.Sprintf("attachment; filename=%s", f.Name)
	c.Response().Header().Set("Content-Disposition", headerVal)
	http.ServeContent(c.Response(), c.Request(), f.Name, time.Now(), f.Inner)
	return nil
}

func (s *Server) Files(c echo.Context) error {
	const op errors.Op = "upload/handler.Files"

	pageParam := c.Param("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	files, _, err := s.uploadRepo.List(page)
	if err != nil {
		return errors.E(err, op)
	}

	return c.JSON(200, files)
}

// Delete deletes specyfied file
func (s *Server) Delete(c echo.Context) error {
	const op errors.Op = "upload/handler.Delete"

	type deletePayload struct {
		IDs []string `json:"ids" validate:"required"`
	}

	p := &deletePayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	if err := c.Validate(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	for _, id := range p.IDs {
		id, err := uuid.FromString(id)
		if err != nil {
			return errors.E(err, errors.Invalid, op)
		}

		err = s.uploadRepo.Delete(id)
		if err != nil {
			return errors.E(err, errors.Invalid, op)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "file deleted successfully"})
}
