package server

import (
	"net/http"
)

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Use(middleware func(http.Handler) http.HandlerFunc) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *Server) Route(pattern string, handler http.HandlerFunc) {
	s.mux.Handle(pattern, applyMiddleware(handler, s.middlewares))
}

func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
