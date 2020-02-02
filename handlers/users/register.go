package users

import (
	"github.com/hichuyamichu-me/uploader/models"
	usersService "github.com/hichuyamichu-me/uploader/services/users"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	inviteID := c.Param("inviteID")
	p := &models.User{}
	if err := c.Bind(p); err != nil {
		return err
	}

	err := usersService.Register(inviteID, p)
	if err != nil {
		return err
	}
	return nil
}
