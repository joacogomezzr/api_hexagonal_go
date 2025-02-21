package routes

import (
	"github.com/gofiber/fiber/v2"
	bookRoutes "api-joaquin/internal/book/routes"
	adminRoutes "api-joaquin/internal/admin/routes"
	bookInterfaces "api-joaquin/internal/book/interface"
	adminInterfaces "api-joaquin/internal/admin/interface"
)

// SetupRoutes configura las rutas de la API
func SetupRoutes(app *fiber.App, bookHandler *bookInterfaces.BookHandler, adminHandler *adminInterfaces.AdminHandler) {
	// Rutas de libros
	bookRoutes.SetupBookRoutes(app, bookHandler)

	// Rutas de administradores
	adminRoutes.SetupAdminRoutes(app, adminHandler)
}
