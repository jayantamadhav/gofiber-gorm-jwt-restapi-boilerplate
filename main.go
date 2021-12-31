package main

import (
	"fmt"
	"gofiber_restapi/database"
	middleware "gofiber_restapi/middleware"
	router "gofiber_restapi/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	app := fiber.New()

	database.ConnectDB()

	middleware.FiberMiddleware(app)

	router.SetupRoutes(app)

	app.Listen(":3000")
}
