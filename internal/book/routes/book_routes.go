package routes

import (
	"github.com/gofiber/fiber/v2"
	"api-joaquin/internal/book/interface"
)

// SetupBookRoutes configura las rutas del recurso libro.
func SetupBookRoutes(app *fiber.App, handler *interfaces.BookHandler) {
	bookGroup := app.Group("/api/v1/books")

	bookGroup.Get("/", handler.GetBooks)
	bookGroup.Post("/register", handler.CreateBook)
	bookGroup.Put("/:id", handler.UpdateBook) 
	bookGroup.Delete("/:id", handler.DeleteBook)
}
