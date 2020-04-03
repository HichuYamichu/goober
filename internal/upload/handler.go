package upload

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hichuyamichu-me/uploader/errors"
	"github.com/hichuyamichu-me/uploader/internal/users"
	"github.com/labstack/echo/v4"
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

	fName := c.Param("name")
	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, fName)
	return c.File(filePath)
}

// FilesInfo handles FilesInfo report
func (h *Handler) FilesInfo(c echo.Context) error {
	const op errors.Op = "upload/handler.FilesInfo"

	files, err := h.uplServ.GetFileData()
	if err != nil {
		return errors.E(err, op)
	}

	type statusResponce struct {
		Name      string `json:"name"`
		Size      int64  `json:"size"`
		CreatedAt string `json:"createdAt"`
		Owner     string `json:"owner"`
	}

	res := make([]*statusResponce, len(files))
	for i, file := range files {
		Uploads := &statusResponce{
			Name:      file.Name(),
			Size:      file.Size(),
			CreatedAt: file.ModTime().Format("2006/01/02 15:04"),
			Owner:     "",
		}
		res[i] = Uploads
	}

	return c.JSON(200, res)
}

// Upload handles file upload
func (h *Handler) Upload(c echo.Context) error {
	const op errors.Op = "upload/handler.Upload"

	form, err := c.MultipartForm()
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}
	files := form.File["files"]

	user := c.Get("user").(*users.User)
	sizeLimit := user.Quota
	sizeTotal := int64(0)
	for _, file := range files {
		sizeTotal += file.Size
		if sizeTotal > int64(sizeLimit) {
			return echo.NewHTTPError(http.StatusRequestEntityTooLarge)
		}
	}

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
		h.uplServ.Save(file)
		u := &uploads{
			URL:  fmt.Sprintf("https://%s/files/%s", domain, file.Filename),
			Name: file.Filename,
			Size: file.Size,
		}
		upl[i] = u
	}

	res := &uploadResult{Success: true, Files: upl}
	return c.JSON(http.StatusOK, res)
}

// Delete deletes specyfied file
func (h *Handler) Delete(c echo.Context) error {
	const op errors.Op = "upload/handler.Delete"

	fName := c.Param("name")
	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, fName)
	err := os.Remove(filePath)
	if err != nil {
		return errors.E(err, errors.IO, op)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "file deleted successfully"})
}
