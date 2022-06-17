package http

import (
	"net/http"

	catalogue "go.topheruk.bookcatalogue"
)

func (s *Service) HandleGetBookReviews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := s.getISBN(w, r)
		s.rr.List(id)
	}
}

func (s *Service) HandleAddBookReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := s.getISBN(w, r)
		var review catalogue.Review
		review.Book = id
		s.rr.Insert(review)

		s.rr.Insert(review)
	}
}
