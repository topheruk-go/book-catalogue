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
	List() ([]Book, error)
	Find(id isbn.ISBN) (*Book, error)
	Add(book Book) error
}
