package main

import (
	"os"

	"github.com/FulgurCode/StellarPing/pkg/mongodb"
	"github.com/FulgurCode/StellarPing/pkg/news"
	"github.com/FulgurCode/StellarPing/server"
	"github.com/FulgurCode/StellarPing/utils"
	"github.com/joho/godotenv"
)

func main() {
	var port string
	godotenv.Load(".env")

	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}

	mongodb.Connect()

	utils.OnceADay(news.GetNews)

	server.Run(port)
}
