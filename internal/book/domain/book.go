//api-joaquin/internal/book/domain/book.go
package domain

// Book representa la entidad de un libro.
type Book struct {
	ID     int64    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
