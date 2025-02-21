
package application

import (
	"api-joaquin/internal/book/domain"
	"api-joaquin/internal/book/domain/repositories"
)

// BookPostUseCase maneja la lógica de creación de un libro.
type BookPostUseCase struct {
	Repo repositories.BookRepository
}

// NewBookPostUseCase inicializa el caso de uso.
func NewBookPostUseCase(repo repositories.BookRepository) *BookPostUseCase {
	return &BookPostUseCase{Repo: repo}
}

// Execute crea un nuevo libro en la base de datos.
func (uc *BookPostUseCase) Execute(book *domain.Book) error {
	return uc.Repo.Create(book)
}
