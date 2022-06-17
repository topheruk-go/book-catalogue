package bookcatalogue

import (
	"github.com/google/uuid"
	"go.topheruk.bookcatalogue/pkg/isbn"
)

type Author struct {
	ID    uuid.UUID   `json:"id"`
	Books []isbn.ISBN `json:"books"`
	Name  string      `json:"name"`
}

type AuthorRepository interface {
	List() ([]Author, error)
	Insert(*Author) error
	Find(uuid.UUID) (*Author, error)
}
