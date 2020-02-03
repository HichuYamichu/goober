package users

import (
	"strconv"

	usersService "github.com/hichuyamichu-me/uploader/services/users"

	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	usersService.Delete(id)
	return nil
}
