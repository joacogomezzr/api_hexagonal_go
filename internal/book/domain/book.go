//api-joaquin/internal/book/domain/book.go
package domain

// Book representa la entidad de un libro.
type Book struct {
<<<<<<< HEAD
	ID     int    `json:"id"`
=======
	ID     int64    `json:"id"`
>>>>>>> 8d61fe1 (c)
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
