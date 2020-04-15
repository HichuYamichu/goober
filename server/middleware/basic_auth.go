package middleware

import (
	"crypto/subtle"
	"fmt"
	"strings"

	"github.com/hichuyamichu-me/goober/errors"
	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	const op errors.Op = "middleware/basic_auth.BasicAuth"

	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		users := viper.GetStringSlice("goober.admin")
		fmt.Println(users)
		for _, user := range users {
			split := strings.Split(user, ":")
			if len(split) != 2 {
				return false, errors.E(op, errors.Internal)
			}
			if subtle.ConstantTimeCompare([]byte(username), []byte(split[0])) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(split[1])) == 1 {
				return true, nil
			}
		}
		return false, nil
	})
}
