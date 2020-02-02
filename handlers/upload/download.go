package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Download(c echo.Context) error {
	fName := c.Param("name")
	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, fName)
	return c.File(filePath)
}
