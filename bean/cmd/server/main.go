package main

import (
	"harvest/bean/internal/adapter/handler"

	"harvest/bean/internal/driver/server"
)

func main() {
	h := handler.New()

	s := server.New()
	s.Route("/login", h.Login)
	s.Route("/", h.Landing)
	s.Listen(":8080")
}
