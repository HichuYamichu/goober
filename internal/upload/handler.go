package upload

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
	fName := c.Param("name")
	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, fName)
	return c.File(filePath)
}

// Status handles status report
func (h *Handler) Status(c echo.Context) error {
	data, err := h.uplServ.generateStatiscics()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(200, data)
}

// Upload handles file upload
func (h *Handler) Upload(c echo.Context) error {
	type uploadResult struct {
		URL     string `json:"url"`
		Name    string `json:"name"`
		Success bool   `json:"success"`
	}

	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	files := form.File["files"]

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sizeLimit := claims["quota"].(float64)
	sizeTotal := int64(0)
	for _, file := range files {
		sizeTotal += file.Size
		if sizeTotal > int64(sizeLimit) {
			return echo.NewHTTPError(http.StatusRequestEntityTooLarge)
		}
	}

	domain := viper.GetString("domain")
	res := make([]*uploadResult, len(files))
	for i, file := range files {
		h.uplServ.save(file)
		r := &uploadResult{
			URL:     fmt.Sprintf("https://%s/api/download/%s", domain, file.Filename),
			Name:    file.Filename,
			Success: true,
		}
		res[i] = r
	}

	return c.JSON(http.StatusOK, res)
}
