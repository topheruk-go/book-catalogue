package psql

import (
	"github.com/google/uuid"
	catalogue "go.topheruk.bookcatalogue"
	// "github.com/jackc/pgx/v4"
)

type AuthorRepository struct {
	db map[uuid.UUID]catalogue.Author
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{db: make(map[uuid.UUID]catalogue.Author)}
}

func (r AuthorRepository) List() ([]catalogue.Author, error) { return nil, catalogue.ErrTodo }

func (r AuthorRepository) Find(id uuid.UUID) (*catalogue.Author, error) {
	return nil, catalogue.ErrTodo
}

func (r *AuthorRepository) Insert(author catalogue.Author) error { return catalogue.ErrTodo }
