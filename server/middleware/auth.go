package middleware

import (
	"bytes"
	"strings"

	"github.com/go-http-utils/headers"
	"github.com/hichuyamichu-me/goober/errors"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/spf13/viper"
)

func Auth() echo.MiddlewareFunc {
	noJWTCheck := !viper.IsSet("jwt")
	if noJWTCheck {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error { return next(c) }
		}
	}

	var verifyFn func(buf []byte) (jwt.Token, error)

	useJWK := viper.IsSet("jwt.jwk_url")
	if useJWK {
		jwkurl := viper.GetString("jwt.jwk_url")

		verifyFn = func(buf []byte) (jwt.Token, error) {
			buf, err := jws.VerifyWithJKU(buf, jwkurl)
			if err != nil {
				return nil, err
			}
			token, err := jwt.Parse(bytes.NewReader(buf))
			if err != nil {
				return nil, err
			}

			return token, nil
		}
	} else {
		key := []byte(viper.GetString("jwt.key"))
		alg := jwa.SignatureAlgorithm(viper.GetString("jwt.alg"))

		verifyFn = func(buf []byte) (jwt.Token, error) {
			token, err := jwt.ParseVerify(bytes.NewReader(buf), alg, key)
			if err != nil {
				return nil, err
			}
			return token, err
		}
	}

	issuer := viper.GetString("jwt.issuer")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get(headers.Authorization)
			splitToken := strings.Split(auth, "Bearer")
			if len(splitToken) != 2 {
				return errors.E(errors.Errorf("invalid jwt"), errors.Authentication)
			}
			auth = strings.TrimSpace(splitToken[1])
			token, err := verifyFn([]byte(auth))
			if err != nil {
				return errors.E(err, errors.Authentication)
			}

			if token.Issuer() != issuer {
				return errors.E(errors.Errorf("invalid issuer"), errors.Authentication)
			}

			claims := token.PrivateClaims()
			c.Set("roles", claims["x-goober-roles"])
			return next(c)
		}
	}
}
