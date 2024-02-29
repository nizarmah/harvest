package server

import (
	"net/http"
)

type Server struct {
	baseHandler BaseHandler
	middlewares []Middleware

	mux *http.ServeMux
}

type Config struct {
	BaseHandler BaseHandler
}

type BaseHandler = func(Handler) func(http.ResponseWriter, *http.Request)

type Handler = func(http.ResponseWriter, *http.Request) error

type Middleware = func(Handler) Handler
