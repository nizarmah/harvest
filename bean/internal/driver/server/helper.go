package server

func applyMiddleware(
	handler Handler,
	middlewares []Middleware,
) Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
