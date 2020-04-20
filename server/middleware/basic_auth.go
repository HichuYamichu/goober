package middleware

import (
	"crypto/subtle"
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	skipper := func(c echo.Context) bool {
		return pathSkip(c.Path()) || !viper.IsSet("admin")
	}

	return middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper: skipper,
		Validator: func(username, password string, c echo.Context) (bool, error) {
			fmt.Println(c.Path())
			users := viper.GetStringSlice("admin")
			for _, user := range users {
				split := strings.Split(user, ":")
				if subtle.ConstantTimeCompare([]byte(username), []byte(split[0])) == 1 &&
					subtle.ConstantTimeCompare([]byte(password), []byte(split[1])) == 1 {
					if len(split) == 3 {
						c.Set("role", split[2])
					}
					return true, nil
				}
			}
			return false, nil
		},
	})
}
