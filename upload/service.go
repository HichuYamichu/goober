package upload

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/hichuyamichu-me/uploader/errors"
	"github.com/spf13/viper"
)

// Service performs operations specyfic to upload domain
type Service struct{}

// NewService creates new upload service
func NewService() *Service {
	return &Service{}
}

// Save saves file to disk
func (s *Service) Save(file *multipart.FileHeader) error {
	const op errors.Op = "upload/service.Save"

	src, err := file.Open()
	if err != nil {
		return errors.E(err, errors.IO, op)
	}
	defer src.Close()

	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return errors.E(err, errors.IO, op)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return errors.E(err, errors.IO, op)
	}
	return nil
}

// GenerateStatiscics generates statistic data
func (s *Service) GenerateStatiscics() ([]os.FileInfo, error) {
	const op errors.Op = "upload/service.GenerateStatiscics"

	uploadDir := viper.GetString("upload_dir")
	files, err := ioutil.ReadDir(uploadDir)
	if err != nil {
		return nil, errors.E(err, errors.IO, op)
	}

	return files, nil
}
