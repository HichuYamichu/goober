package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hichuyamichu-me/goober/errors"
	"github.com/hichuyamichu-me/goober/server/middleware"
	"github.com/hichuyamichu-me/goober/upload"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

// Server main server struct
type Server struct {
	router     *echo.Echo
	uploadRepo *upload.Repository
	db         *gorm.DB
	domain     string
}

// New bootstraps server
func New(db *gorm.DB) (*Server, error) {
	s := &Server{
		router:     echo.New(),
		uploadRepo: upload.NewRepository(db),
		db:         db,
		domain:     viper.GetString("domain"),
	}

	s.router.HideBanner = true
	s.router.HTTPErrorHandler = httpErrorHandler
	s.router.Validator = NewValidator()

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.BodyLimit())

	s.router.Use(middleware.Auth())
	s.router.GET("/:id", s.Download)
	s.router.GET("/list/:page", s.Files)
	s.router.POST("", s.Upload)
	s.router.DELETE("", s.Delete)

	return s, nil
}

func (s *Server) Shutdown(ctx context.Context) {
	_ = s.router.Shutdown(ctx)
	s.db.Close()
}

func (s *Server) Start(host string, port string) error {
	return s.router.Start(fmt.Sprintf("%s:%s", host, port))
}

type uploadFile struct {
	ID     string `json:"id"`
	URL    string `json:"url,omitempty"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Reason error  `json:"reason,omitempty"`
}

func (s *Server) Upload(c echo.Context) error {
	const op errors.Op = "upload/handler.Upload"

	form, err := c.MultipartForm()
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}
	files := form.File["files"]

	type uploadResult struct {
		Succeeded []*uploadFile `json:"succeeded"`
		Failed    []*uploadFile `json:"failed"`
	}

	succeeded := make([]*uploadFile, 0)
	failed := make([]*uploadFile, 0)
	for _, file := range files {
		id, err := s.uploadRepo.Save(file)
		if err != nil {
			u := &uploadFile{Name: file.Filename, Size: file.Size, Reason: err}
			failed = append(failed, u)
		}
		u := &uploadFile{
			ID:   id,
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

	succeeded := make([]*uploadFile, 0)
	failed := make([]*uploadFile, 0)
	for _, id := range p.IDs {
		id, err := uuid.FromString(id)
		if err != nil {
			u := &uploadFile{ID: id.String(), Reason: err}
			failed = append(failed, u)
		}

		file, err := s.uploadRepo.Delete(id)
		if err != nil {
			u := &uploadFile{ID: id.String(), Name: file.Name, Size: file.Size, Reason: err}
			failed = append(failed, u)
		}

		u := &uploadFile{
			ID:   id.String(),
			Name: file.Name,
			Size: file.Size,
		}
		succeeded = append(succeeded, u)
	}

	type deleteResult struct {
		Succeeded []*uploadFile `json:"succeeded"`
		Failed    []*uploadFile `json:"failed"`
	}

	res := &deleteResult{Succeeded: succeeded, Failed: failed}
	return c.JSON(http.StatusOK, res)
}
