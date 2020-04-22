package upload

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hichuyamichu-me/goober/errors"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

// Handler handles all upload domain actions
type Handler struct {
	uplServ *Service
}

// NewHandler creates new Handler
func NewHandler(uplServ *Service) *Handler {
	return &Handler{uplServ: uplServ}
}

// Download handles file downloading
func (h *Handler) Download(c echo.Context) error {
	const op errors.Op = "upload/handler.Download"

	fileID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	file, err := h.uplServ.GetFile(fileID)
	if err != nil {
		return errors.E(err, op)
	}

	f, err := file.Open()
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}
	defer f.Close()

	headerVal := fmt.Sprintf("attachment; filename=%s", file.Name)
	c.Response().Header().Set("Content-Disposition", headerVal)
	http.ServeContent(c.Response(), c.Request(), file.Name, time.Now(), f)
	return nil
}

// Files handles file listing
func (h *Handler) Files(c echo.Context) error {
	const op errors.Op = "upload/handler.Files"

	pageParam := c.Param("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	files, err := h.uplServ.GetFileData(page)
	if err != nil {
		return errors.E(err, op)
	}

	return c.JSON(200, files)
}

// Upload handles file upload
func (h *Handler) Upload(c echo.Context) error {
	const op errors.Op = "upload/handler.Upload"

	form, err := c.MultipartForm()
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}
	files := form.File["files"]

	type uploads struct {
		URL  string `json:"url"`
		Name string `json:"name"`
		Size int64  `json:"size"`
	}

	type uploadResult struct {
		Files   []*uploads `json:"files"`
		Success bool       `json:"success"`
	}

	domain := viper.GetString("domain")
	upl := make([]*uploads, len(files))
	for i, file := range files {
		err := testExtention(file)
		if err != nil {
			return errors.E(op, errors.Invalid, err)
		}
		fName, err := h.uplServ.Save(file)
		if err != nil {
			res := &uploadResult{Success: false, Files: upl}
			return c.JSON(http.StatusInternalServerError, res)
		}
		u := &uploads{
			URL:  fmt.Sprintf("https://%s/files/%s", domain, fName),
			Name: file.Filename,
			Size: file.Size,
		}
		upl[i] = u
	}

	res := &uploadResult{Success: true, Files: upl}
	return c.JSON(http.StatusOK, res)
}

func testExtention(f *multipart.FileHeader) error {
	exts := viper.GetStringSlice("blocked_extentions")
	for _, sufix := range exts {
		if strings.HasSuffix(f.Filename, sufix) {
			return errors.Errorf("illegal file extention")
		}
	}
	return nil
}

// Delete deletes specyfied file
func (h *Handler) Delete(c echo.Context) error {
	const op errors.Op = "upload/handler.Delete"

	type deletePayload struct {
		ID string `json:"id" validate:"required"`
	}

	p := &deletePayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	if err := c.Validate(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	fileID, err := uuid.FromString(p.ID)
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	err = h.uplServ.DeleteFile(fileID)
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "file deleted successfully"})
}
