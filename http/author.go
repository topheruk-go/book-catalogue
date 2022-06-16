package http

import (
	"net/http"

	catalogue "go.topheruk.bookcatalogue"
)

func (s *Service) HandleListAuthors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.sr.ListAuthors()
	}
}

func (s *Service) HandleAddAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author catalogue.Author
		s.sr.AddAuthor(author)
	}
}

func (s *Service) HandleGetAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := s.UUID(w, r)
		s.sr.FindAuthor(id)
	}
}