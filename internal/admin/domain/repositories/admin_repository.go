package repositories

import "api-joaquin/internal/admin/domain"

// AdminRepository define los métodos que debe implementar el repositorio.
type AdminRepository interface {
	Create(admin *domain.Admin) error
	GetAll() ([]domain.Admin, error)
	Update(admin *domain.Admin) error
	Delete(id int) error
}
