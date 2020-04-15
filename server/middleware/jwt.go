package middleware

import (
	"github.com/dgrijalva/jwt-go"
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
