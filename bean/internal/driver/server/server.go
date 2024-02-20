package server

import (
	"net/http"
)

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Route(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
