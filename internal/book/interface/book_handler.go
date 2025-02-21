package interfaces

import (
	"api-joaquin/internal/book/controllers"
	"api-joaquin/internal/book/domain"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// BookHandler maneja las peticiones HTTP relacionadas con libros.
type BookHandler struct {
	Controller *controllers.BookController
}

// NewBookHandler crea un nuevo manejador.
func NewBookHandler(controller *controllers.BookController) *BookHandler {
	return &BookHandler{Controller: controller}
}

// CreateBook maneja la creación de un nuevo libro.
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	book := new(domain.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Solicitud inválida. Verifique los datos enviados.",
			"error":   err.Error(),
		})
	}

	if err := h.Controller.PostUseCase.Execute(book); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo crear el libro.",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Libro creado exitosamente.",
		"data":    book,
	})
}

// GetBooks maneja la obtención de libros.
func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.Controller.GetUseCase.Execute()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudieron obtener los libros.",
			"error":   err.Error(),
		})
	}

	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "success",
			"message": "No se encontraron libros en la base de datos.",
			"data":    books,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Lista de libros obtenida con éxito.",
		"data":    books,
	})
}
// UpdateBook maneja la actualización de un libro.
func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "ID inválido.",
		})
	}

	book := new(domain.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Solicitud inválida. Verifique los datos enviados.",
		})
	}

	book.ID = id
	if err := h.Controller.PutUseCase.Execute(book); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "Libro no encontrado.",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo actualizar el libro.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Libro actualizado correctamente.",
		"data":    book,
	})
}

// DeleteBook maneja la eliminación de un libro.
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "ID inválido.",
		})
	}

	if err := h.Controller.DeleteUseCase.Execute(id); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "Libro no encontrado.",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "No se pudo eliminar el libro.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Libro eliminado correctamente.",
	})
}
