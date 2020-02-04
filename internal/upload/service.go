package upload

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/spf13/viper"
)

type uploadService struct{}

func NewService() *uploadService {
	return &uploadService{}
}

func (s *uploadService) Save(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func (s *uploadService) GenerateStatiscics() ([]string, error) {
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
