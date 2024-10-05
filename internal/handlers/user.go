package handlers

import (
	"net/http"

	"github.com/FulgurCode/StellarPing/pkg/user"
	"github.com/FulgurCode/StellarPing/utils"
	"github.com/FulgurCode/StellarPing/views"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	var component = views.SignUp()

	return utils.Render(c, component)
}

func SignupPost(c echo.Context) error {
	var user user.User
	c.Bind(&user)

	user.Password = utils.BcryptGenerateHash(user.Password)

	if err := user.Signup(); err != nil {
		var component = views.SignUp()

		return utils.Render(c, component)
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", "/login")
		return c.NoContent(200)
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}

func Login(c echo.Context) error {
	var component = views.Login()

	return utils.Render(c, component)
}

func LoginPost(c echo.Context) error {
	var body user.User
	c.Bind(&body)

	var user = user.GetUser(body.Email)

	if correct := utils.BcryptCompare(body.Password, user.Password); !correct {
		return utils.Render(c, views.Login())
	}

	utils.CreateSession(c, "auth")
	utils.AddSessionValue(c, "auth", "email", user.Email)
	utils.AddSessionValue(c, "auth", "isLoggedIn", true)

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", "/")
		return c.NoContent(200)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
