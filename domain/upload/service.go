package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/hichuyamichu-me/goober/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

// Service performs operations specyfic to upload domain
type Service struct {
	fileRepo *Repository
}

// NewService creates new upload service
func NewService(fileRepo *Repository) *Service {
	return &Service{fileRepo: fileRepo}
}

type e struct {
}

func (s *Service) GetFile(id uuid.UUID) (*File, error) {
	const op errors.Op = "upload/service.GetFile"

	file, err := s.fileRepo.FindOne(id)
	if err != nil {
		return nil, errors.E(err, op)
	}
	return file, nil
}

// Save saves file to disk
func (s *Service) Save(file *multipart.FileHeader) (string, error) {
	const op errors.Op = "upload/service.Save"

	src, err := file.Open()
	if err != nil {
		return "", errors.E(err, errors.IO, op)
	}
	defer src.Close()

	fileEntity := &File{Name: file.Filename, Size: file.Size, CreatedAt: time.Now()}
	err = s.fileRepo.Create(fileEntity)
	if err != nil {
		return "", errors.E(err, op)
	}

	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, fileEntity.ID)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", errors.E(err, errors.IO, op)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", errors.E(err, errors.IO, op)
	}
	return fileEntity.ID.String(), nil
}

// GetFileData generates statistic data
func (s *Service) GetFileData(page int) ([]*File, error) {
	const op errors.Op = "upload/service.GetFileData"

	skip := page * 10
	files, err := s.fileRepo.Find(skip)
	if err != nil {
		return nil, errors.E(err, op)
	}

	return files, nil
}

func (s *Service) DeleteFile(id uuid.UUID) error {
	const op errors.Op = "upload/service.DeleteFile"

	file := &File{ID: id}
	err := s.fileRepo.Delete(file)
	if err != nil {
		return errors.E(err, op)
	}

	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, id)
	err = os.Remove(filePath)
	if err != nil {
		return errors.E(err, errors.IO, op)
	}

	return nil
}
