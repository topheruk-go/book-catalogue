package bookcatalogue

import (
	"time"

	"github.com/topheruk-go/isbn"
)

type Book struct {
	Isbn      isbn.ISBN `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

func NewBook(isbn isbn.ISBN, title, author string) *Book {
	return &Book{isbn, title, author, time.Now()}
}

type BookRepository interface {
	List() ([]Book, error)
	Find(id isbn.ISBN) (*Book, error)
	Insert(book *Book) error
}
