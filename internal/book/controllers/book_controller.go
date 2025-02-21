package controllers

import (
	"api-joaquin/internal/book/application"
	"api-joaquin/internal/book/domain/repositories"
)

// inyección de dependencias.
type BookController struct {
	PostUseCase   *application.BookPostUseCase
	GetUseCase    *application.BookGetUseCase
	PutUseCase    *application.BookPutUseCase
	DeleteUseCase *application.BookDeleteUseCase
}

// inicializar el controlador con las dependencias.
func NewBookController(repo repositories.BookRepository) *BookController {
	return &BookController{
		PostUseCase:   application.NewBookPostUseCase(repo),
		GetUseCase:    application.NewBookGetUseCase(repo),
		PutUseCase:    application.NewBookPutUseCase(repo),
		DeleteUseCase: application.NewBookDeleteUseCase(repo),
	}
}
