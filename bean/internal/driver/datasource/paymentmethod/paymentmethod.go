package paymentmethod

import (
	"errors"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/database"
)

type dataSource struct {
	db *database.DB
}

func New(db *database.DB) usecase.PaymentMethodDataSource {
	return &dataSource{db}
}

func (ds *dataSource) Create(paymentMethod *entity.PaymentMethod) (*entity.PaymentMethod, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) FindByUserId(userId int) ([]*entity.PaymentMethod, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) Delete(paymentMethod *entity.PaymentMethod) error {
	return errors.New("not implemented")
}
