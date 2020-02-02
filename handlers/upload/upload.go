package handlers

import (
	"net/http"

	uploadService "github.com/hichuyamichu-me/uploader/services/upload"
	"github.com/labstack/echo/v4"
)

type result struct {
	url  string
	name string
}

func Upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		uploadService.Save(file)
	}

	res := &uploadResponce{
		MSG: "giratka",
	}
	return c.JSON(http.StatusSeeOther, res)
}

type uploadResponce struct {
	MSG string
}
