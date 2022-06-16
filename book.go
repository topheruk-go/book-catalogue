package bookcatalogue

import (
	"time"

	"go.topheruk.bookcatalogue/pkg/isbn"
)

type Book struct {
	ID        isbn.ISBN `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

type BookRepository interface {
	ListBooks() ([]Book, error)
	FindBook(id isbn.ISBN) (*Book, error)
	AddBook(book Book) error
}
