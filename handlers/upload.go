package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

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
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dataDir := c.Get("uploadDir")
		filePath := fmt.Sprintf("%s/%s", dataDir, file.Filename)
		dst, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return c.Redirect(http.StatusSeeOther, "/status")
}
