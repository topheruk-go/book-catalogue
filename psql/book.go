package psql

import (
	catalogue "go.topheruk.bookcatalogue"
	"go.topheruk.bookcatalogue/pkg/isbn"
)

type BookRepository struct {
	db map[isbn.ISBN]catalogue.Book
}

func NewBookRepository() *BookRepository {
	return &BookRepository{db: make(map[isbn.ISBN]catalogue.Book)}
}

func (r BookRepository) ListBooks() ([]catalogue.Book, error) { return nil, catalogue.ErrTodo }

func (r BookRepository) FindBook(id isbn.ISBN) (*catalogue.Book, error) {
	return nil, catalogue.ErrTodo
}

func (r *BookRepository) AddBook(book catalogue.Book) error { return catalogue.ErrTodo }
