package handlers

import (
	"github.com/FulgurCode/StellarPing/utils"
	"github.com/FulgurCode/StellarPing/views"
	"github.com/labstack/echo/v4"
)

// Home page handler
func Home(c echo.Context) error {
	var component = views.Home()

	return utils.Render(c, component)
}
