package main

import (
	handler "harvest/bean/internal/adapter/handler"

	server "harvest/bean/internal/driver/server"
)

func main() {
	h := handler.Init()

	s := server.Init()
	s.Route("/", h.Landing)
	s.Listen(":8080")
}
