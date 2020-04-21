package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-http-utils/headers"
	"github.com/hichuyamichu-me/goober/errors"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/spf13/viper"
)

var keyStore *jwk.Set

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	skip := func(c echo.Context) error { return next(c) }
	noJWTCheck := !viper.IsSet("jwt")
	if noJWTCheck {
		return skip
	}

	useJWK := viper.IsSet("jwt.jwk_url")
	if useJWK {
		var err error
		keyStore, err = jwk.Fetch(viper.GetString("jwt.jwk_url"))
		if err != nil {
			panic(err)
		}
	}

	return func(c echo.Context) error {
		if pathSkip(c.Path()) {
			return next(c)
		}
		return doJWT(c, next)
	}
}

func doJWT(c echo.Context, next echo.HandlerFunc) error {
	auth := c.Request().Header.Get(headers.Authorization)
	splitToken := strings.Split(auth, "Bearer")
	if len(splitToken) != 2 {
		return &echo.HTTPError{
			Code:     http.StatusUnauthorized,
			Message:  "invalid or expired jwt",
			Internal: fmt.Errorf("invalid format"),
		}
	}
	auth = strings.TrimSpace(splitToken[1])

	token := new(jwt.Token)
	token, err := jwt.Parse(auth, keyFunc)

	if err == nil && token.Valid {
		c.Set("user", token)
		claims := token.Claims.(jwt.MapClaims)
		c.Set("role", claims["role"])
		return next(c)
	}
	return &echo.HTTPError{
		Code:     http.StatusUnauthorized,
		Message:  "invalid or expired jwt",
		Internal: err,
	}
}

func keyFunc(t *jwt.Token) (interface{}, error) {
	var alg string
	var key interface{}
	var err error
	if keyStore != nil {
		keys := keyStore.LookupKeyID(t.Header["kid"].(string))
		alg = keys[0].Algorithm()
		key, err = keys[0].Materialize()
		if err != nil {
			return nil, err
		}
	} else {
		alg = viper.GetString("jwt.type")
		key = []byte(viper.GetString("jwt.key"))
	}

	if t.Method.Alg() != alg {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}
	return key, nil
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
