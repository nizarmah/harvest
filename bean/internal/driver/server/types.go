package server

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}
