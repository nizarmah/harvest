package main

import (
	"fmt"

	"harvest/bean/internal/adapter/env"
	"harvest/bean/internal/adapter/handler"

	"harvest/bean/internal/driver/datasource"
	"harvest/bean/internal/driver/server"
)

func main() {
	e, err := env.New()
	if err != nil {
		panic(
			fmt.Errorf("error reading env: %v", err),
		)
	}

	ds, err := datasource.New(&datasource.DsnBuilder{
		Host:        e.DB.Host,
		Name:        e.DB.Name,
		Username:    e.DB.Username,
		Password:    e.DB.Password,
		Tls:         true,
		Interpolate: true,
	})
	if err != nil {
		panic(
			fmt.Errorf("error connecting db: %v", err),
		)
	}
	defer ds.Close()

	h := handler.New()

	s := server.New()

	s.Route("/login", h.Login)
	s.Route("/", h.Landing)

	fmt.Println("server started on :8080")
	s.Listen(":8080")
}
