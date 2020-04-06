package upload

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

// File representrs file
type File struct {
	ID        uuid.UUID `gorm:"type:uuid;" json:"id"`
	Name      string    `gorm:"not null;" json:"name"`
	Size      int64     `gorm:"not null;" json:"size"`
	CreatedAt int64     `gorm:"not null;" json:"createdAt"`
}

func (file *File) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

func (file *File) Open() (*os.File, error) {
	uploadDir := viper.GetString("upload_dir")
	filePath := fmt.Sprintf("%s/%s", uploadDir, file.ID)
	return os.Open(filePath)
}
