package upload

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"

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

type fileData struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
}

// GenerateStatiscics generates statistic data
func (s *Service) GenerateStatiscics() ([]*fileData, error) {
	uploadDir := viper.GetString("upload_dir")
	files, err := ioutil.ReadDir(uploadDir)
	if err != nil {
		return nil, err
	}

	res := make([]*fileData, len(files))
	for i, file := range files {
		fileData := &fileData{
			Name:      file.Name(),
			Size:      file.Size(),
			CreatedAt: file.ModTime(),
		}
		res[i] = fileData
	}

	return res, nil
}
