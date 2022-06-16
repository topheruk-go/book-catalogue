package http

import (
	"net/http"

	cv5 "github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	catalogue "go.topheruk.bookcatalogue"
	"go.topheruk.bookcatalogue/pkg/isbn"
)

type Service struct {
	br catalogue.BookRepository
	rr catalogue.ReviewRepository
	sr catalogue.AuthorRepository
	m  *cv5.Mux
}

func NewService(br catalogue.BookRepository, rr catalogue.ReviewRepository, sr catalogue.AuthorRepository) *Service {
	s := &Service{br, rr, sr, cv5.NewMux()}
	s.routes()
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.m.ServeHTTP(w, r) }

func (s *Service) routes() {
	s.m.Get("/book", s.HandleListBooks())
	s.m.Get("/book/{isbn}", s.HandleGetBook())
	s.m.Post("/book", s.HandleAddBook())

	s.m.Get("/book/{isbn}/review/", s.HandleGetBookReviews())
	s.m.Post("/book/{isbn}/review", s.HandleAddBookReview())

	s.m.Get("/author", s.HandleListAuthors())
	s.m.Get("/author/{uuid}", s.HandleGetAuthor())
	s.m.Post("/author", s.HandleAddAuthor())
}

func (s *Service) ISBN(w http.ResponseWriter, r *http.Request) (isbn.ISBN, error) {
	id := cv5.URLParam(r, "isbn")
	return isbn.Parse(id)
}

func (s *Service) UUID(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	id := cv5.URLParam(r, "uuid")
	return uuid.Parse(id)
}
