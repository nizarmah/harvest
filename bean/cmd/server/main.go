package main

import (
	"fmt"

	envAdapter "harvest/bean/internal/adapter/env"
	"harvest/bean/internal/entity"

	"harvest/bean/internal/driver/database"
	userDS "harvest/bean/internal/driver/datasource/user"
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

	users := userDS.New(db)

	u, err := users.FindByEmail("test@gmail.com")
	if u != nil {
		fmt.Println("user found by email: ", u.ID, u.Email, u.CreatedAt, u.UpdatedAt)
	} else {
		fmt.Println("user not found by email: ", err)
	}

	u, err = users.Create(&entity.User{Email: "test@gmail.com"})
	if u != nil {
		fmt.Println("user created: ", u.ID, u.Email, u.CreatedAt, u.UpdatedAt)
	} else {
		fmt.Println("user not created: ", err)
		return
	}

	u, err = users.FindById(u.ID)
	if u != nil {
		fmt.Println("user found by id: ", u.ID, u.Email, u.CreatedAt, u.UpdatedAt)
	} else {
		fmt.Println("user not found by id: ", err)
	}
}
