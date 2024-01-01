package token

import (
	"errors"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/database"
)

type dataSource struct {
	db *database.DB
}

func New(db *database.DB) usecase.LoginTokenDataSource {
	return &dataSource{
		db: db,
	}
}

func (ds *dataSource) Create(inputToken *entity.LoginToken) (*entity.LoginToken, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) FindByEmail(userId int) (*entity.LoginToken, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) Delete(token *entity.LoginToken) (*entity.LoginToken, error) {
	return nil, errors.New("not implemented")
}
