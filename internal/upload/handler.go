package upload

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type uploadHandler struct {
	uplServ *uploadService
}

func NewHandler(uplServ *uploadService) *uploadHandler {
	return &uploadHandler{uplServ: uplServ}
}

func (h *uploadHandler) Download(c echo.Context) error {
	fName := c.Param("name")
	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, fName)
	return c.File(filePath)
}

func (h *uploadHandler) Status(c echo.Context) error {
	data, err := h.uplServ.GenerateStatiscics()
	if err != nil {
		return err
	}
	return c.JSON(200, data)
}

func (h *uploadHandler) Upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sizeLimit := claims["quota"].(float64)
	if sizeLimit == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing upload permission",
		})
	}

	size := int64(0)
	for _, file := range files {
		size += file.Size
		if size > int64(sizeLimit) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "payload too large",
			})
		}
	}

	domain := viper.GetString("domain")
	res := make([]*uploadResult, len(files))
	for i, file := range files {
		h.uplServ.Save(file)
		r := &uploadResult{
			URL:     fmt.Sprintf("https://%s/api/download/%s", domain, file.Filename),
			Name:    file.Filename,
			Success: true,
		}
		res[i] = r
	}

	return c.JSON(http.StatusOK, res)
}
