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

func New(db *database.DB) usecase.TokenDataSource {
	return &dataSource{
		db: db,
	}
}

func (ds *dataSource) Create(inputToken *entity.Token) (*entity.Token, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) FindByUserId(userId int) (*entity.Token, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) FindByHashedToken(hashedToken string) (*entity.Token, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) Delete(token *entity.Token) (*entity.Token, error) {
	return nil, errors.New("not implemented")
}
