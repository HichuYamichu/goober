package middleware

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
	"github.com/hichuyamichu-me/goober/errors"
	"github.com/labstack/echo/v4"
)

func CanRead(next echo.HandlerFunc) echo.HandlerFunc {
	skip := func(c echo.Context) error { return next(c) }
	noAuth := !viper.IsSet("jwt") && !viper.IsSet("admin")
	noRoleCheck := !viper.IsSet("roles")
	if noAuth || noRoleCheck {
		return skip
	}

	return func(c echo.Context) error {
		if pathSkip(c.Path()) {
			return next(c)
		}

		return canRead(c, next)
	}
}

func canRead(c echo.Context, next echo.HandlerFunc) error {
	const op errors.Op = "middleware/perms.canRead"

	userPerms := getUserPerms(c)
	canRead := strings.Contains(userPerms, "r")
	if !canRead {
		return errors.E(errors.Authentication, op)
	}

	if err := next(c); err != nil {
		errors.E(err, op)
	}
	return nil
}

func CanWrite(next echo.HandlerFunc) echo.HandlerFunc {
	skip := func(c echo.Context) error { return next(c) }
	noAuth := !viper.IsSet("jwt") && !viper.IsSet("admin")
	noRoleCheck := !viper.IsSet("roles")
	if noAuth || noRoleCheck {
		return skip
	}

	return func(c echo.Context) error {
		if pathSkip(c.Path()) {
			return next(c)
		}

		return canWrite(c, next)
	}
}

func canWrite(c echo.Context, next echo.HandlerFunc) error {
	const op errors.Op = "middleware/perms.canWrite"

	userPerms := getUserPerms(c)
	canWrite := strings.Contains(userPerms, "w")
	if !canWrite {
		return errors.E(op, errors.Authentication)
	}

	if err := next(c); err != nil {
		errors.E(err, op)
	}
	return nil
}

func CanDelete(next echo.HandlerFunc) echo.HandlerFunc {

	skip := func(c echo.Context) error { return next(c) }
	noAuth := !viper.IsSet("jwt") && !viper.IsSet("admin")
	noRoleCheck := !viper.IsSet("roles")
	if noAuth || noRoleCheck {
		return skip
	}

	return func(c echo.Context) error {
		if pathSkip(c.Path()) {
			return next(c)
		}

		return canDelete(c, next)
	}
}

func canDelete(c echo.Context, next echo.HandlerFunc) error {
	const op errors.Op = "middleware/perms.canDelete"

	userPerms := getUserPerms(c)
	canDelete := strings.Contains(userPerms, "d")
	if !canDelete {
		return errors.E(op, errors.Authentication)
	}

	if err := next(c); err != nil {
		errors.E(err, op)
	}
	return nil
}

func getUserPerms(c echo.Context) string {
	userRole, ok := c.Get("role").(string)
	if !ok {
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return ""
		}
		claims := user.Claims.(jwt.MapClaims)
		userRole = claims["role"].(string)
	}

	var userRolePerms string
	for _, configRole := range viper.GetStringSlice("roles") {
		split := strings.Split(configRole, ":")
		configRoleName := split[0]
		configRolePerms := split[1]
		if configRoleName == userRole {
			userRolePerms = configRolePerms
		}
	}

	return userRolePerms
}
