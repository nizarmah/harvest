package server

import (
	"net/http"
)

func Init() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Route(path string, handler HandlerFunc) {
	s.mux.Handle(path, handler)
}

func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
