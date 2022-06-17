package http

import (
	"net/http"

	"github.com/topheruk-go/isbn"
	catalogue "go.topheruk.bookcatalogue"
)

// GET /book/{isbn}/review
func (s *Service) HandleGetBookReviews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := s.parseISBN(w, r)
		if err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}

		reveiws, err := s.rr.ListByIndex(id)
		if err != nil {
			s.respond(w, r, err, http.StatusNotFound)
			return
		}

		s.respond(w, r, reveiws, http.StatusOK)
	}
}

// POST /book/{isbn}/review
func (s *Service) HandleAddBookReview() http.HandlerFunc {
	type dto struct {
		Book  isbn.ISBN `json:"book"`
		Name  string    `json:"name"`
		Score int       `json:"score"`
		Text  string    `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := s.parseISBN(w, r)
		if err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}

		var v dto
		if err := s.decode(w, r, &v); err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}

		review := catalogue.NewReview(id, v.Name, v.Text, v.Score)
		if err := s.rr.Insert(review); err != nil {
			s.respond(w, r, err, http.StatusInternalServerError)
			return
		}

		s.created(w, r, review.ID.String())
	}
}
