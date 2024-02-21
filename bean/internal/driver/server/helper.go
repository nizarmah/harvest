package server

import (
	"net/http"
)

func applyMiddleware(
	handler http.Handler,
	middlewares []func(http.Handler) http.HandlerFunc,
) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
