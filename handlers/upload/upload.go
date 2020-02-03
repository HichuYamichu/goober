package handlers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	uploadService "github.com/hichuyamichu-me/uploader/services/upload"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type uploadResult struct {
	URL     string `json:"url"`
	Name    string `json:"name"`
	Success bool   `json:"success"`
}

func Upload(c echo.Context) error {
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
		uploadService.Save(file)
		r := &uploadResult{
			URL:     fmt.Sprintf("https://%s/api/download/%s", domain, file.Filename),
			Name:    file.Filename,
			Success: true,
		}
		res[i] = r
	}

	return c.JSON(http.StatusOK, res)
}
