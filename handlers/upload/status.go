package handlers

import (
	uploadService "github.com/hichuyamichu-me/uploader/services/upload"
	"github.com/labstack/echo/v4"
)

func Status(c echo.Context) error {
	data, err := uploadService.GenerateStatiscics()
	if err != nil {
		return err
	}
	return c.JSON(200, data)
}
