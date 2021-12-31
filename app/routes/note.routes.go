package routes

import (
	handlers "gofiber_restapi/app/handlers"
	middleware "gofiber_restapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	note.Post("/", handlers.CreateNote)
	note.Get("/", middleware.JWTProtected(), handlers.GetNotes)
	note.Get("/:noteId", middleware.JWTProtected(), handlers.GetNote)
	note.Put("/:noteId", middleware.JWTProtected(), handlers.UpdateNote)
	note.Delete("/:noteId", middleware.JWTProtected(), handlers.DeleteNote)
}
