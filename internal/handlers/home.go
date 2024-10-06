package handlers

import (
	"net/http"

	"github.com/FulgurCode/StellarPing/pkg/news"
	"github.com/FulgurCode/StellarPing/utils"
	"github.com/FulgurCode/StellarPing/views"
	"github.com/labstack/echo/v4"
)

// Home page handler
func Home(c echo.Context) error {
	// if isLoggedIn := utils.GetSessionValue(c, "auth", "isLoggedIn"); isLoggedIn == nil || !isLoggedIn.(bool) {
	// 	if c.Request().Header.Get("HX-Request") == "true" {
	// 		c.Response().Header().Set("HX-Location", "/login")
	// 		return c.NoContent(200)
	// 	}

	// 	return c.Redirect(http.StatusSeeOther, "/login")
	// }

	var news = news.GetNews()

	var component = views.Home(news)
	return utils.Render(c, component)
}

func News(c echo.Context) error {
	// if isLoggedIn := utils.GetSessionValue(c, "auth", "isLoggedIn"); isLoggedIn == nil || !isLoggedIn.(bool) {
	// 	if c.Request().Header.Get("HX-Request") == "true" {
	// 		c.Response().Header().Set("HX-Location", "/login")
	// 		return c.NoContent(200)
	// 	}

	// 	return c.Redirect(http.StatusSeeOther, "/login")
	// }
	var id = c.Param("id")
	var n = news.GetOneNews(id)

	var component = views.News(n)
	return utils.Render(c, component)
}
