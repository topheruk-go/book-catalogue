package http

import (
	"net/http"

	cv5 "github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	ptth "github.com/topheruk-go/util/http"
	bc "go.topheruk.bookcatalogue"
	"go.topheruk.bookcatalogue/pkg/isbn"
)

type Service struct {
	br bc.BookRepository
	rr bc.ReviewRepository
	ar bc.AuthorRepository
	m  *cv5.Mux
}

func NewService(br bc.BookRepository, rr bc.ReviewRepository, ar bc.AuthorRepository) *Service {
	s := &Service{br, rr, ar, cv5.NewMux()}
	s.routes()
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.m.ServeHTTP(w, r) }

func (s *Service) parseISBN(w http.ResponseWriter, r *http.Request) (isbn.ISBN, error) {
	id := cv5.URLParam(r, "isbn")
	return isbn.Parse(id)
}

func (s *Service) parseUUID(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	id := cv5.URLParam(r, "uuid")
	return uuid.Parse(id)
}

func (s Service) created(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Add("Location", "//"+r.Host+r.URL.Path+"/"+id)
	s.respond(w, r, nil, http.StatusCreated)
}

func (s Service) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	ptth.Respond(w, r, data, status)
}

func (s Service) decode(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return ptth.Decode(w, r, data)
}
