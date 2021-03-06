package bookcatalogue

import (
	"time"

	"github.com/google/uuid"
	"github.com/topheruk-go/isbn"
)

type Review struct {
	ID        uuid.UUID `json:"id"`
	Book      isbn.ISBN `json:"book"`
	Name      string    `json:"name"`
	Text      string    `json:"text"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReview(book isbn.ISBN, name, text string, score int) *Review {
	return &Review{uuid.New(), book, name, text, score, time.Now()}
}

type ReviewRepository interface {
	Insert(*Review) error
	ListByIndex(isbn.ISBN) ([]Review, error)
}
