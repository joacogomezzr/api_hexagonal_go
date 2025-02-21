package interfaces

import (
	"api-joaquin/internal/admin/controllers"
	"api-joaquin/internal/admin/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AdminHandler maneja las peticiones HTTP relacionadas con administradores.
type AdminHandler struct {
	Controller *controllers.AdminController
}

// NewAdminHandler crea un nuevo manejador.
func NewAdminHandler(controller *controllers.AdminController) *AdminHandler {
	return &AdminHandler{Controller: controller}
}

// CreateAdmin maneja el registro de un nuevo administrador.
func (h *AdminHandler) CreateAdmin(c *fiber.Ctx) error {
	admin := new(domain.Admin)
	if err := c.BodyParser(admin); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Solicitud inválida. Verifique los datos enviados.",
			"error":   err.Error(),
		})
	}

	if err := h.Controller.PostUseCase.Execute(admin); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo registrar el administrador.",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Administrador registrado exitosamente.",
		"data":    admin,
	})
}

// GetAdmins maneja la obtención de administradores.
func (h *AdminHandler) GetAdmins(c *fiber.Ctx) error {
	admins, err := h.Controller.GetUseCase.Execute()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudieron obtener los administradores.",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Lista de administradores obtenida con éxito.",
		"data":    admins,
	})
}


func (h *AdminHandler) UpdateAdmin(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID inválido."})
	}

	admin := new(domain.Admin)
	if err := c.BodyParser(admin); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Solicitud inválida. Verifique los datos enviados.", "error": err.Error()})
	}

	admin.ID = id
	if err := h.Controller.PutUseCase.Execute(admin); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "No se pudo actualizar el administrador.", "error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Administrador actualizado correctamente.", "data": admin})
}

// DeleteAdmin maneja la eliminación de un administrador.
func (h *AdminHandler) DeleteAdmin(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID inválido."})
	}

	if err := h.Controller.DeleteUseCase.Execute(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "No se pudo eliminar el administrador."})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Administrador eliminado correctamente."})
}
