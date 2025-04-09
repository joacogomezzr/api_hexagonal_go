package infrastructure

import (
	"api-joaquin/internal/admin/domain"
	"api-joaquin/internal/admin/domain/repositories"
	"database/sql"
	"log"
)

type AdminService struct {
	DB *sql.DB
}

func NewAdminService(db *sql.DB) repositories.AdminRepository {
	return &AdminService{DB: db}
}

func (s *AdminService) Create(admin *domain.Admin) error {
	query := "INSERT INTO admins (name, email, password) VALUES (?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Println("Error preparando la consulta:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Name, admin.Email, admin.Password)
	
	result, err := stmt.Exec(admin.Name, admin.Email, admin.Password)
	id, err := result.LastInsertId()	
	if err != nil {
		log.Println("Error ejecutando la consulta:", err)
		return err
	}
	admin.ID = int(id)

	return nil
}


func (s *AdminService) GetAll() ([]domain.Admin, error) {
	query := "SELECT id, name, email FROM admins"
	rows, err := s.DB.Query(query)
	if err != nil {
		log.Println("Error consultando los administradores:", err)
		return nil, err
	}
	defer rows.Close()

	var admins []domain.Admin
	for rows.Next() {
		var admin domain.Admin
		if err := rows.Scan(&admin.ID, &admin.Name, &admin.Email); err != nil {
			log.Println("Error escaneando fila:", err)
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func (s *AdminService) Update(admin *domain.Admin) error {
	query := "UPDATE admins SET name = ?, email = ? WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Println("Error preparando la actualización:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Name, admin.Email, admin.ID)
	return err
}

func (s *AdminService) Delete(id int) error {
	query := "DELETE FROM admins WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Println(" Error preparando la eliminación:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}