package server

import (
	"log"
	"os"

	"github.com/FulgurCode/StellarPing/internal/handlers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(port string) {
	var app = echo.New()

	// Remove trailing slash to avoid 404 errors
	app.Pre(middleware.RemoveTrailingSlash())

	// Creating session
	var sessionSecret = os.Getenv("SESSION_SECRET")
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionSecret))))

	// add static files
	app.Static("/static", "assets")

	app.GET("/", handlers.Home)
	app.GET("/news/:id", handlers.News)

	app.GET("/signup", handlers.SignUp)
	app.POST("/signup", handlers.SignupPost)

	app.GET("/login", handlers.Login)
	app.POST("/login", handlers.LoginPost)

	log.Fatal(app.Start(":" + port))
}
