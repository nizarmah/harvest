package main

import (
	"fmt"

	envAdapter "harvest/bean/internal/adapter/env"
	"harvest/bean/internal/adapter/handler"

	"harvest/bean/internal/driver/postgres"
	"harvest/bean/internal/driver/server"
	"harvest/bean/internal/driver/template"
	"harvest/bean/internal/driver/view/landing"
	"harvest/bean/internal/driver/view/login"
)

func main() {
	env, err := envAdapter.New()
	if err != nil {
		panic(
			fmt.Errorf("error reading env: %v", err),
		)
	}

	db, err := postgres.New(&postgres.DSNBuilder{
		Host:     env.DB.Host,
		Port:     env.DB.Port,
		Name:     env.DB.Name,
		Username: env.DB.Username,
		Password: env.DB.Password,
		SSLMode:  "disable",
	})
	if err != nil {
		panic(
			fmt.Errorf("error connecting db: %v", err),
		)
	}
	defer db.Close()

	landingView, err := landing.New(template.FS, template.LandingTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating landing view: %v", err),
		)
	}

	loginView, err := login.New(template.FS, template.LoginTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating login view: %v", err),
		)
	}

	h := handler.New(landingView, loginView)

	s := server.New()

	s.Route("/", h.Landing)
	s.Route("/get-started", h.Login)

	s.Listen(":8080")
}
