package main

import (
	"fmt"

	envAdapter "harvest/bean/internal/adapter/env"

	"harvest/bean/internal/driver/database"
)

func main() {
	env, err := envAdapter.New()
	if err != nil {
		panic(
			fmt.Errorf("error reading env: %v", err),
		)
	}

	db, err := database.New(&database.DSNBuilder{
		Host:        env.DB.Host,
		Name:        env.DB.Name,
		Username:    env.DB.Username,
		Password:    env.DB.Password,
		Tls:         true,
		Interpolate: true,
		ParseTime:   true,
	})
	if err != nil {
		panic(
			fmt.Errorf("error connecting db: %v", err),
		)
	}
	defer db.Close()
}
