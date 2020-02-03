package users

import (
	usersService "github.com/hichuyamichu-me/uploader/services/users"
	"github.com/labstack/echo/v4"
)

func Invite(c echo.Context) error {
	conf := &usersService.UserConfig{}
	if err := c.Bind(c); err != nil {
		return err
	}
	id := usersService.GenereateInvite(conf)
	return c.String(200, id)
}
