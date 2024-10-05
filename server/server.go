package server

import (
	"log"

	"github.com/FulgurCode/StellarPing/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(port string) {
	var app = echo.New()

	// Remove trailing slash to avoid 404 errors
	app.Pre(middleware.RemoveTrailingSlash())

	// add static files
	app.Static("/static", "assets")

	app.GET("/", handlers.Home)

	app.GET("/signup", handlers.SignUp)
	app.POST("/signup", handlers.SignupPost)
	app.GET("/login", handlers.Login)


	log.Fatal(app.Start(":" + port))
}
