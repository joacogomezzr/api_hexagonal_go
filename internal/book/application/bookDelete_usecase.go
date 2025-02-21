package application

import "api-joaquin/internal/book/domain/repositories"

// BookDeleteUseCase maneja la eliminación de libros.
type BookDeleteUseCase struct {
	Repo repositories.BookRepository
}

// NewBookDeleteUseCase inicializa el caso de uso.
func NewBookDeleteUseCase(repo repositories.BookRepository) *BookDeleteUseCase {
	return &BookDeleteUseCase{Repo: repo}
}

// Execute elimina un libro por su ID.
func (uc *BookDeleteUseCase) Execute(id int) error {
	return uc.Repo.Delete(id)
}
