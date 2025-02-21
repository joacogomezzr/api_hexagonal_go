//api-joaqui/pkg/middleware/middleware.go
package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// SetupCORS configura los encabezados CORS
func SetupCORS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Manejo de preflight request
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		
		return c.Next()
	}
}
