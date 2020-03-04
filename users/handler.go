package users

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hichuyamichu-me/uploader/errors"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// Handler handles all user domain actions
type Handler struct {
	usrServ *Service
}

// NewHandler creates new Handler
func NewHandler(usrServ *Service) *Handler {
	h := &Handler{usrServ: usrServ}
	return h
}

// CreateUser handles user creation
func (h Handler) CreateUser(c echo.Context) error {
	const op errors.Op = "users/handler.CreateUser"

	user := &User{}
	if err := c.Bind(user); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	err := h.usrServ.CreateUser(user)
	if err != nil {
		return errors.E(err, op)
	}
	return c.JSON(200, map[string]interface{}{"message": "user created successfuly"})
}

// UpdateUser handles user updates
func (h Handler) UpdateUser(c echo.Context) error {
	const op errors.Op = "users/handler.UpdateUser"

	return nil
}

// DeleteUser handles deleting the user
func (h Handler) DeleteUser(c echo.Context) error {
	const op errors.Op = "users/handler.DeleteUser"

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}
	h.usrServ.DeleteUser(id)
	return c.JSON(200, map[string]interface{}{"message": "user deleted successfuly"})
}

// ChangePass handles password change
func (h *Handler) ChangePass(c echo.Context) error {
	const op errors.Op = "users/handler.ChangePass"

	type passChangePayload struct {
		Pass string `json:"password"`
	}

	p := &passChangePayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	fmt.Println(userID)

	err := h.usrServ.ChangePassword(int(userID), p.Pass)
	if err != nil {
		return errors.E(err, op)
	}
	return nil
}

// Login handles user login
func (h Handler) Login(c echo.Context) error {
	const op errors.Op = "users/handler.Login"

	type loginPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	p := &loginPayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	user, err := h.usrServ.VerifyCredentials(p.Username, p.Password)
	if err != nil {
		return errors.E(err, op)
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
		Token string `json:"token"`
		User  *User  `json:"user"`
	}

	res := &loginResponce{
		Token: t,
		User:  user,
	}
	return c.JSON(http.StatusOK, res)
}
