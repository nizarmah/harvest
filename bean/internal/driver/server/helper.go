package server

import (
	"net/http"
)

func applyMiddleware(
	handler http.Handler,
	middlewares []func(http.Handler) http.HandlerFunc,
) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
