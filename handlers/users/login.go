package users

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	usersService "github.com/hichuyamichu-me/uploader/services/users"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type safeUser struct {
	Username string `json:"username"`
	Quota    int64  `json:"quota"`
	Admin    bool   `json:"admin"`
}

type loginResponce struct {
	Token string    `json:"token"`
	User  *safeUser `json:"user"`
}

func Login(c echo.Context) error {
	p := &loginPayload{}
	if err := c.Bind(p); err != nil {
		return err
	}

	user := usersService.FindOneByUsername(p.Username)
	if user == nil {
		return echo.ErrUnauthorized
	}

	match := usersService.CheckPasswordHash(p.Password, user.Pass)
	if !match {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["username"] = user.Username
	claims["admin"] = user.Admin
	claims["quota"] = user.Quota

	t, err := token.SignedString([]byte(viper.GetString("secret_key")))
	if err != nil {
		return err
	}

	res := &loginResponce{
		Token: t,
		User: &safeUser{
			Username: user.Username,
			Quota:    user.Quota,
			Admin:    user.Admin,
		},
	}
	return c.JSON(http.StatusOK, res)
}
