package server

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

type Handler = http.Handler
type HandlerFunc = http.HandlerFunc
