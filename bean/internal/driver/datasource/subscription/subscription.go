package subscription

import (
	"errors"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/database"
)

type dataSource struct {
	db *database.DB
}

func New(db *database.DB) usecase.SubscriptionDataSource {
	return &dataSource{
		db: db,
	}
}

func (ds *dataSource) Create(sub *entity.Subscription) (*entity.Subscription, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) FindByUserId(userId string) ([]*entity.Subscription, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) Update(sub *entity.Subscription) (*entity.Subscription, error) {
	return nil, errors.New("not implemented")
}

func (ds *dataSource) Delete(sub *entity.Subscription) error {
	return errors.New("not implemented")
}
