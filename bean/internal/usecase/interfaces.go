package usecase

import (
	"harvest/bean/internal/entity"
)

type SubscriptionDataSource interface {
	Create(subscription *entity.Subscription) (*entity.Subscription, error)

	FindByUserId(userId int) ([]*entity.Subscription, error)

	Delete(subscription *entity.Subscription) error
}

type UserDataSource interface {
	Create(user *entity.User) (*entity.User, error)

	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type LoginTokenDataSource interface {
	Create(token *entity.LoginToken) error

	FindUnexpired(token *entity.LoginToken) (*entity.LoginToken, error)

	Delete(token *entity.LoginToken) error
}
