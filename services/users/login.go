package users

import (
	"github.com/hichuyamichu-me/uploader/models"
	usersRepo "github.com/hichuyamichu-me/uploader/store/users"
	"github.com/labstack/echo"
)

func Login(username, password string) error {
	user := &models.User{Username: username}
	user = usersRepo.FindOne(user)
	if user == nil {
		return echo.ErrUnauthorized
	}

	match := CheckPasswordHash(password, user.Pass)
	if !match {
		return echo.ErrUnauthorized
	}
	return nil
}
