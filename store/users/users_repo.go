package users

import (
	"github.com/hichuyamichu-me/uploader/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InjectDB(dbInstance *gorm.DB) {
	db = dbInstance
}

func FindOne(where *models.User) *models.User {
	user := &models.User{}
	if err := db.Where(where).First(&user).Error; err != nil {
		return nil
	}
	return user
}

func Create(user *models.User) {
	db.Create(user)
}
