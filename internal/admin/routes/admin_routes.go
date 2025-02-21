package routes

import (
	"github.com/gofiber/fiber/v2"
	"api-joaquin/internal/admin/interface"
)

// SetupAdminRoutes configura las rutas del recurso administrador.
func SetupAdminRoutes(app *fiber.App, handler *interfaces.AdminHandler) {
	adminGroup := app.Group("/api/v1/admins")

	adminGroup.Post("/register", handler.CreateAdmin)
	adminGroup.Get("/", handler.GetAdmins)
	adminGroup.Put("/:id", handler.UpdateAdmin)
	adminGroup.Delete("/:id", handler.DeleteAdmin)
}
