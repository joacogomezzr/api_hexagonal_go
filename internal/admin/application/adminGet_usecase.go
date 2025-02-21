package application

import (
	"api-joaquin/internal/admin/domain"
	"api-joaquin/internal/admin/domain/repositories"
)

// AdminGetUseCase maneja la lógica para obtener administradores.
type AdminGetUseCase struct {
	Repo repositories.AdminRepository
}

// NewAdminGetUseCase inicializa el caso de uso.
func NewAdminGetUseCase(repo repositories.AdminRepository) *AdminGetUseCase {
	return &AdminGetUseCase{Repo: repo}
}

// Execute obtiene todos los administradores.
func (uc *AdminGetUseCase) Execute() ([]domain.Admin, error) {
	return uc.Repo.GetAll()
}
