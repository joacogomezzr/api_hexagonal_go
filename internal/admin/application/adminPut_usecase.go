package application

import (
	"api-joaquin/internal/admin/domain"
	"api-joaquin/internal/admin/domain/repositories"
)

type AdminPutUseCase struct {
	Repo repositories.AdminRepository
}

func NewAdminPutUseCase(repo repositories.AdminRepository) *AdminPutUseCase {
	return &AdminPutUseCase{Repo: repo}
}

func (uc *AdminPutUseCase) Execute(admin *domain.Admin) error {
	return uc.Repo.Update(admin)
}
