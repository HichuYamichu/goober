package users

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/hichuyamichu-me/uploader/errors"
	"github.com/labstack/echo/v4"
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

// ChangePass handles password change
func (h *Handler) ChangePass(c echo.Context) error {
	const op errors.Op = "users/handler.ChangePass"

	type passChangePayload struct {
		Pass string `json:"password" validate:"required"`
	}

	p := &passChangePayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	if err := c.Validate(p); err != nil {
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

// ActivateUser handles activation of user account
func (h *Handler) ActivateUser(c echo.Context) error {
	const op errors.Op = "auth/handler.ActivateUser"

	type activateUserPayload struct {
		ID int `json:"id" validate:"required"`
	}

	p := &activateUserPayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	if err := c.Validate(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	err := h.usrServ.ActivateUser(p.ID)
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "user activated"})
}
