package user

import (
	"errors"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/database"
)

type dataSource struct {
	db *database.DB
}

func New(db *database.DB) usecase.UserDataSource {
	return &dataSource{
		db: db,
	}
}

func (ds *dataSource) Create(inputUser *entity.User) (*entity.User, error) {
	res, err := ds.db.Pool.Exec("INSERT INTO users (email) VALUES (?)", inputUser.Email)
	if err != nil {
		return nil, errors.New("error inserting user")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.New("error getting last insert id")
	}

	user := &entity.User{}

	err = ds.db.Pool.
		QueryRow("SELECT * FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, errors.New("error getting inserted user")
	}

	return user, nil
}

func (ds *dataSource) FindById(id int) (*entity.User, error) {
	user := &entity.User{}

	err := ds.db.Pool.
		QueryRow("SELECT * FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, errors.New("error getting user")
	}

	return user, nil
}

func (ds *dataSource) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}

	err := ds.db.Pool.
		QueryRow("SELECT * FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, errors.New("error getting user")
	}

	return user, nil
}
