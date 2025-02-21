package application

import (
	"api-joaquin/internal/book/domain"
	"api-joaquin/internal/book/domain/repositories"
)

// BookGetUseCase maneja la lógica de obtener libros.
type BookGetUseCase struct {
	Repo repositories.BookRepository
}

// NewBookGetUseCase inicializa el caso de uso.
func NewBookGetUseCase(repo repositories.BookRepository) *BookGetUseCase {
	return &BookGetUseCase{Repo: repo}
}

// Execute obtiene todos los libros.
func (uc *BookGetUseCase) Execute() ([]domain.Book, error) {
	books, err := uc.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}
