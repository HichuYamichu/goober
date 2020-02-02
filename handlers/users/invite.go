package users

import (
	usersService "github.com/hichuyamichu-me/uploader/services/users"
	"github.com/labstack/echo/v4"
)

func Invite(c echo.Context) error {
	p := &usersService.Permissions{}
	if err := c.Bind(p); err != nil {
		return err
	}
	id := usersService.GenereateInvite(p)
	return c.String(200, id)
}
