package application

import "api-joaquin/internal/admin/domain/repositories"


type AdminDeleteUseCase struct {
	Repo repositories.AdminRepository
}


func NewAdminDeleteUseCase(repo repositories.AdminRepository) *AdminDeleteUseCase {
	return &AdminDeleteUseCase{Repo: repo}
}


func (uc *AdminDeleteUseCase) Execute(id int) error {
	return uc.Repo.Delete(id)
}
