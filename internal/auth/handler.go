package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hichuyamichu-me/uploader/errors"
	"github.com/hichuyamichu-me/uploader/internal/users"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// Handler handles all auth domain actions
type Handler struct {
	authSrv *Service
	usrServ *users.Service
}

// NewHandler creates new Handler
func NewHandler(authSrv *Service, usrServ *users.Service) *Handler {
	h := &Handler{authSrv: authSrv, usrServ: usrServ}
	return h
}

// Login handles user login
func (h *Handler) Login(c echo.Context) error {
	const op errors.Op = "auth/handler.Login"

	type loginPayload struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	p := &loginPayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	if err := c.Validate(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	user, err := h.authSrv.VerifyCredentials(p.Username, p.Password)
	if err != nil {
		return errors.E(err, op)
	}

	if !user.Active {
		return errors.E(errors.Errorf("user not active"), errors.Authentication, op)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["admin"] = user.Admin
	claims["quota"] = user.Quota

	t, err := token.SignedString([]byte(viper.GetString("secret_key")))
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}

	type loginResponce struct {
		Token string      `json:"token"`
		User  *users.User `json:"user"`
	}

	res := &loginResponce{
		Token: t,
		User:  user,
	}
	return c.JSON(http.StatusOK, res)
}

// Register handles user register
func (h *Handler) Register(c echo.Context) error {
	const op errors.Op = "auth/handler.Register"

	type registerPayload struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	p := &registerPayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	if err := c.Validate(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	err := h.usrServ.CreateUser(p.Username, p.Password)
	if err != nil {
		return errors.E(err, op)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "register request waiting for approval"})
}
