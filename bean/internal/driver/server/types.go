package server

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

type HandlerFunc = http.HandlerFunc
