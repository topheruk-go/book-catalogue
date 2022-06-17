package http

import (
	"net/http"

	catalogue "go.topheruk.bookcatalogue"
)

// GET /author
//
// TODO -- URL Params
func (s *Service) HandleListAuthors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := s.ar.List()
		if err != nil {
			s.respond(w, r, err, http.StatusInternalServerError)
			return
		}

		s.respond(w, r, authors, http.StatusOK)
	}
}

// POST /author
func (s *Service) HandleAddAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author catalogue.Author
		if err := s.decode(w, r, &author); err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}

		if err := s.ar.Insert(&author); err != nil {
			s.respond(w, r, err, http.StatusInternalServerError)
			return
		}

		s.created(w, r, author.ID.String())
	}
}

// GET /author/{uuid}
func (s *Service) HandleGetAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := s.parseUUID(w, r)
		if err != nil {
			// Error is a wrong format, should this be bad request?
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}

		author, err := s.ar.Find(id)
		if err != nil {
			s.respond(w, r, err, http.StatusNotFound)
			return
		}

		s.respond(w, r, author, http.StatusOK)
	}
}
