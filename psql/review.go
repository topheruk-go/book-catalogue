package psql

import (
	"github.com/google/uuid"
	catalogue "go.topheruk.bookcatalogue"
	"go.topheruk.bookcatalogue/pkg/isbn"
)

type ReviewRepository struct {
	db map[uuid.UUID]catalogue.Review
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{db: make(map[uuid.UUID]catalogue.Review)}
}

func (db *ReviewRepository) ListBookReviews(isbn isbn.ISBN) ([]catalogue.Review, error) {
	return nil, catalogue.ErrTodo
}

func (db *ReviewRepository) AddBookReview(review catalogue.Review) ([]catalogue.Review, error) {
	return nil, catalogue.ErrTodo
}
