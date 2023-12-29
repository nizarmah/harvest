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

type TokenRepository interface {
	Create(token *entity.Token) (*entity.Token, error)

	FindByUserId(userId string) (*entity.Token, error)

	Delete(token *entity.Token) (*entity.Token, error)
}
