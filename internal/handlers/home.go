package handlers

import (
	"net/http"

	"github.com/FulgurCode/StellarPing/utils"
	"github.com/FulgurCode/StellarPing/views"
	"github.com/labstack/echo/v4"
)

// Home page handler
func Home(c echo.Context) error {
	var component = views.Home()
	if isLoggedIn  := utils.GetSessionValue(c,"auth","isLoggedIn"); isLoggedIn != nil && isLoggedIn.(bool) {
		
		return utils.Render(c, component)
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Location", "/login")
		return c.NoContent(200)
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}
