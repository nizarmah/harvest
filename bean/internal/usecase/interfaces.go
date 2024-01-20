package usecase

import (
	"harvest/bean/internal/entity"
)

type SubscriptionDataSource interface {
	Create(subscription *entity.Subscription) (*entity.Subscription, error)

	FindByUserId(userId int) ([]*entity.Subscription, error)

	Delete(subscription *entity.Subscription) error
}

type PaymentMethodDataSource interface {
	Create(paymentMethod *entity.PaymentMethod) (*entity.PaymentMethod, error)

	FindByUserId(userId int) ([]*entity.PaymentMethod, error)

	Delete(paymentMethod *entity.PaymentMethod) error
}

type UserDataSource interface {
	Create(user *entity.User) (*entity.User, error)

	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type LoginTokenDataSource interface {
	Create(email string, hashedToken string) (*entity.LoginToken, error)

	FindUnexpired(id string) (*entity.LoginToken, error)

	Delete(id string) (*entity.LoginToken, error)
}
