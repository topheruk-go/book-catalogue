package http

import (
	"net/http"

	catalogue "go.topheruk.bookcatalogue"
)

func (s *Service) HandleGetBookReviews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := s.ISBN(w, r)
		s.rr.ListBookReviews(id)
	}
}

func (s *Service) HandleAddBookReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := s.ISBN(w, r)
		var review catalogue.Review
		review.Book = id
		s.rr.AddBookReview(review)

		s.rr.AddBookReview(review)
	}
}
