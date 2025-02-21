package repositories

import "api-joaquin/internal/book/domain"

// BookRepository define los métodos que debe implementar el repositorio.
type BookRepository interface {
	Create(book *domain.Book) error
	GetAll() ([]domain.Book, error)
	Update(book *domain.Book) error
	Delete(id int) error
}
