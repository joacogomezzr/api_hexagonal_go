package repositories

import "api-joaquin/internal/book/domain"
type RabbitMQRepository interface {
	CreateMessageRabbit(messageBook *domain.Book) error
}