package users

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type userHandler struct {
	usrServ *userService
}

func NewHandler(usrServ *userService) *userHandler {
	h := &userHandler{usrServ: usrServ}
	return h
}

func (h userHandler) DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	h.usrServ.DeleteUser(id)
	return nil
}

func (h userHandler) Invite(c echo.Context) error {
	conf := &UserConfig{}
	if err := c.Bind(c); err != nil {
		return err
	}
	id := h.usrServ.GenereateInvite(conf)
	return c.String(200, id)
}

func (h userHandler) Login(c echo.Context) error {
	p := &loginPayload{}
	if err := c.Bind(p); err != nil {
		return err
	}

	user, err := h.usrServ.VerifyCredentials(p.Username, p.Password)
	if err != nil {
		return err
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

func (h userHandler) Register(c echo.Context) error {
	inviteID := c.Param("inviteID")
	p := &User{}
	if err := c.Bind(p); err != nil {
		return err
	}

	err := h.usrServ.CreateUserFromInvite(inviteID, p)
	if err != nil {
		return err
	}
	return nil
}

func (h userHandler) UpdateUser(c echo.Context) error {
	return nil
}
