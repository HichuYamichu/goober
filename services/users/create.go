package users

import (
	"github.com/hichuyamichu-me/uploader/models"
	usersRepo "github.com/hichuyamichu-me/uploader/store/users"
)

func Create(user *models.User) {
	usersRepo.Create(user)
}
