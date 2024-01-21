package usecase

import (
	"harvest/bean/internal/entity"
)

type SubscriptionDataSource interface {
	Create(subscription *entity.Subscription) (*entity.Subscription, error)

	FindByUserId(userId string) ([]*entity.Subscription, error)

	Delete(subscription *entity.Subscription) error
}

type PaymentMethodDataSource interface {
	Create(paymentMethod *entity.PaymentMethod) (*entity.PaymentMethod, error)

	FindByUserId(userId string) ([]*entity.PaymentMethod, error)

	Delete(paymentMethod *entity.PaymentMethod) error
}

type UserDataSource interface {
	Create(email string) (*entity.User, error)

	FindById(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)

	Delete(id string) error
}

type LoginTokenDataSource interface {
	Create(email string, hashedToken string) (*entity.LoginToken, error)

	FindUnexpired(id string) (*entity.LoginToken, error)

	Delete(id string) error
}
