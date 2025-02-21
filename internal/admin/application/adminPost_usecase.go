package application

import (
	"api-joaquin/internal/admin/domain"
	"api-joaquin/internal/admin/domain/repositories"
)

//  creación de un administrador.
type AdminPostUseCase struct {
	Repo repositories.AdminRepository
}

// inicializa  caso de uso.
func NewAdminPostUseCase(repo repositories.AdminRepository) *AdminPostUseCase {
	return &AdminPostUseCase{Repo: repo}
}

// registra un nuevo administrador en la base de datos.
func (uc *AdminPostUseCase) Execute(admin *domain.Admin) error {
	return uc.Repo.Create(admin)
}
