package handlers

import (
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

type statusResponce struct {
	files []string
}

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

	res := &statusResponce{files: fileNames}
	return c.JSON(200, res)
}
