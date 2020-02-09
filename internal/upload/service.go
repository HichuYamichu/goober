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

type fileData struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
	Owner     string    `json:"owner"`
}

func (s *uploadService) GenerateStatiscics() ([]*fileData, error) {
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
			Owner:     "",
		}
		res[i] = fileData
	}

	return res, nil
}
