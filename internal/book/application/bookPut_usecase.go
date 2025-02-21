package application

import (
	"api-joaquin/internal/book/domain"
	"api-joaquin/internal/book/domain/repositories"
)

// BookPutUseCase maneja la actualización de libros.
type BookPutUseCase struct {
	Repo repositories.BookRepository
}

// NewBookPutUseCase inicializa el caso de uso.
func NewBookPutUseCase(repo repositories.BookRepository) *BookPutUseCase {
	return &BookPutUseCase{Repo: repo}
}

// Execute actualiza un libro en la base de datos.
func (uc *BookPutUseCase) Execute(book *domain.Book) error {
	return uc.Repo.Update(book)
}
