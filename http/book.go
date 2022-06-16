package http

import (
	"net/http"

	catalogue "go.topheruk.bookcatalogue"
)

// GET /book/
func (s *Service) HandleListBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.br.ListBooks()
	}
}

// GET /book/{isbn}
func (s *Service) HandleGetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := s.ISBN(w, r)
		s.br.FindBook(id)
	}
}

// POST /book
func (s *Service) HandleAddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book catalogue.Book

		s.br.AddBook(book)
	}
}
