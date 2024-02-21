package server

import (
	"net/http"
)

type Server struct {
	middlewares []func(http.Handler) http.HandlerFunc

	mux *http.ServeMux
}
