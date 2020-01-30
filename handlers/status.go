package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Status(c echo.Context) error {
	dataDir := c.Get("uploadDir")
	files, err := ioutil.ReadDir(dataDir.(string))
	if err != nil {
		return err
	}

	fileNames := make([]string, len(files))
	for _, file := range files {
		fName := file.Name()
		fileNames = append(fileNames, fName)
	}

	return c.Render(http.StatusOK, "status", fileNames)
}
