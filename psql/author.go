package psql

import (
	"github.com/google/uuid"
	catalogue "go.topheruk.bookcatalogue"
)

type AuthorRepository struct {
	db map[uuid.UUID]catalogue.Author
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{db: make(map[uuid.UUID]catalogue.Author)}
}

func (r AuthorRepository) ListAuthors() ([]catalogue.Author, error) { return nil, catalogue.ErrTodo }

func (r AuthorRepository) FindAuthor(id uuid.UUID) (*catalogue.Author, error) {
	return nil, catalogue.ErrTodo
}

func (r *AuthorRepository) AddAuthor(author catalogue.Author) error { return catalogue.ErrTodo }
