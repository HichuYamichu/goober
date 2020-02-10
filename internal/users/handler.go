package users

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	user := &User{}
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := h.usrServ.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(200, map[string]interface{}{"message": "user created successfuly"})
}

// UpdateUser handles user updates
func (h Handler) UpdateUser(c echo.Context) error {
	return nil
}

// DeleteUser handles deleting the user
func (h Handler) DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	h.usrServ.DeleteUser(id)
	return c.JSON(200, map[string]interface{}{"message": "user deleted successfuly"})
}

// ChangePass handles password change
func (h *Handler) ChangePass(c echo.Context) error {
	type passChangePayload struct {
		Pass string `json:"password"`
	}

	p := &passChangePayload{}
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	fmt.Println(userID)

	err := h.usrServ.ChangePassword(int(userID), p.Pass)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return nil
}

// Login handles user login
func (h Handler) Login(c echo.Context) error {
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

	p := &loginPayload{}
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, err := h.usrServ.VerifyCredentials(p.Username, p.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
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
		return echo.NewHTTPError(http.StatusInternalServerError)
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
