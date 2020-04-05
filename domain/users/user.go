package users

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User representrs user
type User struct {
	ID       uuid.UUID `gorm:"type:uuid;" json:"id"`
	Username string    `gorm:"unique;not null;" json:"username"`
	Pass     string    `gorm:"not null;" json:"-"`
	Admin    bool      `gorm:"not null;" json:"admin"`
	Active   bool      `gorm:"not null;" json:"active"`
	Token    string    `gorm:"not null;" json:"token"`
	Quota    int64     `gorm:"not null;" json:"quota"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
