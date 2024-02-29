package server

import (
	"net/http"
)

func New(cfg *Config) *Server {
	return &Server{
		baseHandler: cfg.BaseHandler,

		mux: http.NewServeMux(),
	}
}

func (s *Server) Use(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *Server) Route(pattern string, handler Handler) {
	s.mux.HandleFunc(
		pattern,
		s.baseHandler(
			applyMiddleware(
				handler,
				s.middlewares,
			),
		),
	)
}

func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
