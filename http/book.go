package http

import (
	"net/http"

	"github.com/topheruk-go/isbn"
	catalogue "go.topheruk.bookcatalogue"
)

// GET /book/
func (s *Service) HandleListBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.br.List()
	}
}

// GET /book/{isbn}
func (s *Service) HandleGetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := s.parseISBN(w, r)
		s.br.Find(id)
	}
}

// POST /book
func (s *Service) HandleAddBook() http.HandlerFunc {
	type dto struct {
		ID     isbn.ISBN `json:"id"`
		Title  string    `json:"title"`
		Author string    `json:"author"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var v dto
		if err := s.decode(w, r, &v); err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}

		book := catalogue.NewBook(v.ID, v.Title, v.Author)
		if err := s.br.Insert(book); err != nil {
			s.respond(w, r, err, http.StatusInternalServerError)
			return
		}

		s.created(w, r, book.Isbn.String())
	}
}
