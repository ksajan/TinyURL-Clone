package main

import (
	"os"
	routes "tinyURL/services/api"
	logging "tinyURL/utils"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func initRoutes(app *fiber.App) {
	app.Post("/shorten", routes.ShortenURL)
	app.Get("/:shortURLEncodedValue", routes.ResolveURL)
}

func main() {

	logger := logging.NewLogger()
	err := godotenv.Load()
	if err != nil {
		logger.Errorln("Error loading .env file")
	}
	logger.Infoln("Hello, playground")
	app := fiber.New()
	app.Use(fiberLogger.New())
	initRoutes(app)
	logger.Infoln("Listening on post", os.Getenv("APP_PORT"))
	logger.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
