package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Aadmin(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Admins test page.")

}