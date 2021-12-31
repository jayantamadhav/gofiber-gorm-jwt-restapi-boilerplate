package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	routes "gofiber_restapi/app/routes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	routes.AuthRoutes(api)
	routes.NoteRoutes(api)
}
