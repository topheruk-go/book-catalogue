package bookcatalogue

import (
	"time"

	"github.com/google/uuid"
	"go.topheruk.bookcatalogue/pkg/isbn"
)

type Review struct {
	ID        uuid.UUID `json:"id"`
	Book      isbn.ISBN `json:"book"`
	Name      string    `json:"name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewRepository interface {
	Insert(Review) error
	List(isbn.ISBN) ([]Review, error)
}
