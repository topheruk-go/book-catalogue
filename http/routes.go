package http

func (s *Service) routes() {
	s.m.Get("/book", s.HandleListBooks())
	s.m.Get("/book/{isbn}", s.HandleGetBook())
	s.m.Post("/book", s.HandleAddBook())

	s.m.Get("/book/{isbn}/review", s.HandleGetBookReviews())
	s.m.Post("/book/{isbn}/review", s.HandleAddBookReview())

	s.m.Get("/author", s.HandleListAuthors())
	s.m.Get("/author/{uuid}", s.HandleGetAuthor())
	s.m.Post("/author", s.HandleAddAuthor())
}
