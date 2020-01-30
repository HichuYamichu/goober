package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Download(c echo.Context) error {
	fName := c.Param("name")
	dataDir := c.Get("uploadDir")
	filePath := fmt.Sprintf("%s/%s", dataDir, fName)
	return c.File(filePath)
}
