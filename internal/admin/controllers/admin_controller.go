package controllers

import (
	"api-joaquin/internal/admin/application"
	"api-joaquin/internal/admin/domain/repositories"
)

type AdminController struct {
	PostUseCase   *application.AdminPostUseCase
	GetUseCase    *application.AdminGetUseCase
	PutUseCase    *application.AdminPutUseCase
	DeleteUseCase *application.AdminDeleteUseCase
}

func NewAdminController(repo repositories.AdminRepository) *AdminController {
	return &AdminController{
		PostUseCase:   application.NewAdminPostUseCase(repo),
		GetUseCase:    application.NewAdminGetUseCase(repo),
		PutUseCase:    application.NewAdminPutUseCase(repo),
		DeleteUseCase: application.NewAdminDeleteUseCase(repo),
	}
}
