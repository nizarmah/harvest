package user

import (
	"context"
	"fmt"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/postgres"
)

type dataSource struct {
	db *postgres.DB
}

func New(db *postgres.DB) usecase.UserDataSource {
	return &dataSource{
		db: db,
	}
}

func (ds *dataSource) Create(email string) (*entity.User, error) {
	user := &entity.User{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			("INSERT INTO users (email)"+
				" VALUES ($1)"+
				" RETURNING *"),
			email,
		).
		Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (ds *dataSource) FindById(id string) (*entity.User, error) {
	user := &entity.User{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			"SELECT * FROM users WHERE id = $1",
			id,
		).
		Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	return user, nil
}

func (ds *dataSource) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			"SELECT * FROM users WHERE email = $1",
			email,
		).
		Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return user, nil
}

func (ds *dataSource) Delete(id string) error {
	_, err := ds.db.Pool.
		Exec(
			context.Background(),
			"DELETE FROM users WHERE id = $1",
			id,
		)

	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
