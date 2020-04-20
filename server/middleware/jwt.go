package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hichuyamichu-me/goober/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func JWT() echo.MiddlewareFunc {
	skipper := func(c echo.Context) bool {
		return pathSkip(c.Path()) || !viper.IsSet("jwt")
	}

	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:       skipper,
		SigningKey:    []byte(viper.GetString("jwt.key")),
		SigningMethod: viper.GetString("jwt.type"),
	})
}

func ISS(next echo.HandlerFunc) echo.HandlerFunc {
	skip := func(c echo.Context) error { return next(c) }
	noJWTCheck := !viper.IsSet("jwt")
	noISSCheck := !viper.IsSet("jwt.issuer")
	if noJWTCheck || noISSCheck {
		return skip
	}

	return func(c echo.Context) error {
		if pathSkip(c.Path()) {
			return next(c)
		}
		return iss(c, next)
	}
}

func iss(c echo.Context, next echo.HandlerFunc) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	allowedIssuer := viper.GetString("jwt.issuer")
	match := claims.VerifyIssuer(allowedIssuer, true)
	if !match {
		return errors.E(errors.Errorf("invalid issuer"), errors.Authentication)
	}
	if err := next(c); err != nil {
		return errors.E(err)
	}
	return nil
}
