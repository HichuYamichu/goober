package middleware

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
	"github.com/hichuyamichu-me/goober/errors"
	"github.com/labstack/echo/v4"
)

func ParsePermissions(next echo.HandlerFunc) echo.HandlerFunc {
	if !viper.IsSet("jwt") || !viper.IsSet("jwt.roles") {
		return func(c echo.Context) error { return next(c) }
	}

	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		userRoles := claims.Roles

		for _, userRole := range userRoles {
			role := viper.GetString(fmt.Sprintf("roles.%s", userRole))
			canRead := strings.Contains(role, "r")
			if canRead {
				c.Set("read", true)
			}

			canWrite := strings.Contains(role, "w")
			if canWrite {
				c.Set("write", true)
			}

			canDelete := strings.Contains(role, "d")
			if canDelete {
				c.Set("delete", true)
			}
		}

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func CanRead(next echo.HandlerFunc) echo.HandlerFunc {
	const op errors.Op = "middleware/perms.CanRead"

	if !viper.IsSet("jwt") || !viper.IsSet("jwt.roles") {
		return func(c echo.Context) error { return next(c) }
	}

	return func(c echo.Context) error {
		canRead := c.Get("read").(bool)
		if !canRead {
			return errors.E(op, errors.Authentication)
		}

		if err := next(c); err != nil {
			errors.E(err, op)
		}
		return nil
	}
}

func CanWrite(next echo.HandlerFunc) echo.HandlerFunc {
	const op errors.Op = "middleware/perms.CanWrite"

	if !viper.IsSet("jwt") || !viper.IsSet("jwt.roles") {
		return func(c echo.Context) error { return next(c) }
	}

	return func(c echo.Context) error {
		canWrite := c.Get("write").(bool)
		if !canWrite {
			return errors.E(op, errors.Authentication)
		}

		if err := next(c); err != nil {
			errors.E(err, op)
		}
		return nil
	}
}

func CanDelete(next echo.HandlerFunc) echo.HandlerFunc {
	const op errors.Op = "middleware/perms.CanDelete"

	if !viper.IsSet("jwt") || !viper.IsSet("jwt.roles") {
		return func(c echo.Context) error { return next(c) }
	}

	return func(c echo.Context) error {
		canDelete := c.Get("delete").(bool)
		if !canDelete {
			return errors.E(op, errors.Authentication)
		}

		if err := next(c); err != nil {
			errors.E(err, op)
		}
		return nil
	}
}
