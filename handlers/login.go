package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/")
}

func ServeLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}
