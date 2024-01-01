package usecase

import (
	"harvest/bean/internal/entity"
)

type SubscriptionRepository interface {
	Create(subscription *entity.Subscription) (*entity.Subscription, error)

	FindByUserId(userId string) (*entity.Subscription, error)

	Update(subscription *entity.Subscription) (*entity.Subscription, error)

	Delete(subscription *entity.Subscription) (*entity.Subscription, error)
}

type UserDataSource interface {
	Create(user *entity.User) (*entity.User, error)

	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type LoginTokenDataSource interface {
	Create(token *entity.LoginToken) (*entity.LoginToken, error)

	FindByEmail(userId int) (*entity.LoginToken, error)

	Delete(token *entity.LoginToken) (*entity.LoginToken, error)
}
