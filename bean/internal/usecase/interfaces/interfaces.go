package interfaces

import (
	"harvest/bean/internal/entity"
)

// --- Data Sources ---

type SubscriptionDataSource interface {
	Create(
		userID string,
		paymentMethodID string,
		label string,
		provider string,
		amount int,
		interval int,
		period entity.SubscriptionPeriod,
	) (*entity.Subscription, error)

	FindByID(userID string, id string) (*entity.SubscriptionWithPaymentMethod, error)
	FindByUserID(userID string) ([]*entity.SubscriptionWithPaymentMethod, error)

	Delete(userID string, id string) error
}

type PaymentMethodDataSource interface {
	Create(
		userID string,
		label string,
		last4 string,
		brand entity.PaymentMethodBrand,
		expMonth int,
		expYear int,
	) (*entity.PaymentMethod, error)

	FindByID(userID string, id string) (*entity.PaymentMethod, error)
	FindByUserID(userID string) ([]*entity.PaymentMethod, error)

	Delete(userID string, id string) error
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

// --- Misc ---

type Hasher interface {
	Hash(string) (string, error)
	Compare(string, string) error
}

type Emailer interface {
	Send(email string, subject string, body string) error
}
