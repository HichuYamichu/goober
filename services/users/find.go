package users

import (
	"github.com/hichuyamichu-me/uploader/models"
	usersRepo "github.com/hichuyamichu-me/uploader/store/users"
)

func FindOneByUsername(username string) *models.User {
	user := &models.User{Username: username}
	user = usersRepo.FindOne(user)
	return user
}
