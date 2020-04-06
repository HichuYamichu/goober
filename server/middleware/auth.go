package middleware

import (
	"github.com/hichuyamichu-me/goober/domain/users"
	"github.com/hichuyamichu-me/goober/errors"
	"github.com/labstack/echo/v4"
)

func (mws *Service) LoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	const op errors.Op = "middleware/auth.LoggedIn"

	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		user, err := mws.usersRepo.FindByToken(token)
		if err != nil {
			return errors.E(err, errors.Authentication, op)
		}

		if user == nil {
			return errors.E(errors.Errorf("unauthorized"), errors.Authentication, op)
		}

		c.Set("user", user)

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}

func (mws *Service) Admin(next echo.HandlerFunc) echo.HandlerFunc {
	const op errors.Op = "middleware/auth.Admin"

	return func(c echo.Context) error {
		userVal := c.Get("user")
		user, ok := userVal.(*users.User)
		if !ok {
			return errors.E(errors.Errorf("unauthorized"), errors.Authentication, op)
		}

		if !user.Admin {
			return errors.E(errors.Errorf("unauthorized"), errors.Authentication, op)
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
