package routes

import (
	handlers "gofiber_restapi/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	// auth.Get("/get-token", handlers.GetToken)
	auth.Post("/get-token", handlers.GetToken)
}
