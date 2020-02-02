package upload

import (
	"io/ioutil"

	"github.com/spf13/viper"
)

func GenerateStatiscics() ([]string, error) {
	uploadDir := viper.GetString("upload_dir")
	files, err := ioutil.ReadDir(uploadDir)
	if err != nil {
		return nil, err
	}

	fileNames := make([]string, len(files))
	for _, file := range files {
		fName := file.Name()
		fileNames = append(fileNames, fName)
	}

	return fileNames, nil
}
