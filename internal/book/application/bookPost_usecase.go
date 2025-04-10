package application

import (
	"api-joaquin/internal/book/domain"
	"api-joaquin/internal/book/domain/repositories"
)

// BookPostUseCase maneja la lógica de creación de un libro.
type BookPostUseCase struct {
	Repo repositories.BookRepository
	rabbit repositories.RabbitMQRepository
}

// NewBookPostUseCase inicializa el caso de uso.
func NewBookPostUseCase(repo repositories.BookRepository, rabbitmq repositories.RabbitMQRepository) *BookPostUseCase {
	return &BookPostUseCase{
		Repo: repo,
		rabbit: rabbitmq,
	}
}

// Execute crea un nuevo libro en la base de datos.
func (uc *BookPostUseCase) Execute(book *domain.Book) error {
    err := uc.Repo.Create(book)
    if err != nil {
        return err
    }

    err = uc.rabbit.CreateMessageRabbit(book)
    if err != nil {
        return err
    }

    return nil
}