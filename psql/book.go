package psql

import (
	"github.com/jackc/pgx/v4"
	catalogue "go.topheruk.bookcatalogue"
	"go.topheruk.bookcatalogue/pkg/isbn"
)

type BookRepository struct {
	c  *pgx.Conn
	db map[isbn.ISBN]catalogue.Book
}

func NewBookRepository(conn *pgx.Conn) *BookRepository {
	return &BookRepository{c: conn}
}

func (r BookRepository) List() ([]catalogue.Book, error) { return nil, catalogue.ErrTodo }

func (r BookRepository) Find(id isbn.ISBN) (*catalogue.Book, error) {
	return nil, catalogue.ErrTodo
}

func (r *BookRepository) Insert(book catalogue.Book) error { return catalogue.ErrTodo }
