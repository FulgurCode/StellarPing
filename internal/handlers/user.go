package handlers

import (
	"net/http"

	"github.com/FulgurCode/StellarPing/pkg/user"
	"github.com/FulgurCode/StellarPing/utils"
	"github.com/labstack/echo/v4"
)

func SignupPost(c echo.Context) error {
	var user user.User
	c.Bind(&user)

	user.Password = utils.BcryptGenerateHash(user.Password)

	if err := user.Signup(); err != nil {
		if c.Request().Header.Get("HX-Request") == "true" {
			c.Response().Header().Set("HX-Location", "/signup")
			return c.NoContent(200)
		}

		return c.Redirect(http.StatusSeeOther, "/signup")
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", "/login")
		return c.NoContent(200)
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}
