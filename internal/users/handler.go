package users

import (
	"net/http"
	"strconv"

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
		Password string `json:"password" validate:"required"`
	}

	p := &passChangePayload{}
	if err := c.Bind(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	if err := c.Validate(p); err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	user := c.Get("user").(*User)

	err := h.usrServ.ChangePassword(user.ID, p.Password)
	if err != nil {
		return errors.E(err, op)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "password changed"})
}

// ActivateUser handles activation of user account
func (h *Handler) ActivateUser(c echo.Context) error {
	const op errors.Op = "auth/handler.ActivateUser"

	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	err = h.usrServ.ActivateUser(userID)
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "user activated"})
}

// ListUsers handles user listing
func (h *Handler) ListUsers(c echo.Context) error {
	const op errors.Op = "auth/handler.ListUsers"

	users, err := h.usrServ.ListUsers()
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}

	return c.JSON(http.StatusOK, users)
}

// DeleteUser handles user deletion
func (h *Handler) DeleteUser(c echo.Context) error {
	const op errors.Op = "auth/handler.DeleteUser"

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.E(err, errors.Invalid, op)
	}

	err = h.usrServ.DeleteUser(id)
	if err != nil {
		return errors.E(err, errors.Internal, op)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "user deleted successfully"})
}
