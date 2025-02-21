package infrastructure

import (
	"api-joaquin/internal/book/domain"
	"api-joaquin/internal/book/domain/repositories"
	"database/sql"
	"log"
)

// BookService implementa la interfaz BookRepository.
type BookService struct {
	DB *sql.DB
}

// NewBookService crea una nueva instancia del servicio.
func NewBookService(db *sql.DB) repositories.BookRepository {
	return &BookService{DB: db}
}

// Create agrega un nuevo libro a la base de datos.
func (s *BookService) Create(book *domain.Book) error {
	query := "INSERT INTO books (title, author, year) VALUES (?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Println("❌ Error preparando la consulta:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author, book.Year)
	if err != nil {
		log.Println("❌ Error ejecutando la consulta:", err)
		return err
	}

	return nil
}

// GetAll obtiene todos los libros de la base de datos.
func (s *BookService) GetAll() ([]domain.Book, error) {
	query := "SELECT id, title, author, year FROM books"
	rows, err := s.DB.Query(query)
	if err != nil {
		log.Println("❌ Error consultando los libros:", err)
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			log.Println("❌ Error escaneando fila:", err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
// Update actualiza un libro en la base de datos.
func (s *BookService) Update(book *domain.Book) error {
	query := "UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Println("❌ Error preparando la actualización:", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		log.Println("❌ Error ejecutando la actualización:", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("⚠ No se encontró el libro con el ID especificado.")
		return sql.ErrNoRows
	}

	return nil
}

// Delete elimina un libro por su ID.
func (s *BookService) Delete(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Println("❌ Error preparando la eliminación:", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Println("❌ Error ejecutando la eliminación:", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("⚠ No se encontró el libro con el ID especificado.")
		return sql.ErrNoRows
	}

	return nil
}