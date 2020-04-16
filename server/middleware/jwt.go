package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hichuyamichu-me/goober/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type jwtCustomClaims struct {
	Roles []string `json:"x-goober-roles"`
	jwt.StandardClaims
}

func JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(viper.GetString("goober.jwt.key")),
		SigningMethod: viper.GetString("goober.jwt.type"),
		Claims:        &jwtCustomClaims{},
	})
}

func ISS(next echo.HandlerFunc) echo.HandlerFunc {
	const op errors.Op = "middleware/jwt.ISS"

	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		allowedIssuer := viper.GetString("goober.jwt.issuer")
		if allowedIssuer != "" {
			if claims.Issuer != allowedIssuer {
				return errors.E(errors.Errorf("invalid issuer"), errors.Authentication, op)
			}
		}

		if err := next(c); err != nil {
			return errors.E(err, op)
		}
		return nil
	}
}
