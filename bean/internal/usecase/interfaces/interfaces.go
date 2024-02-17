package interfaces

import (
	"github.com/whatis277/harvest/bean/internal/entity/model"
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
		period model.SubscriptionPeriod,
	) (*model.Subscription, error)

	FindByID(userID string, id string) (*model.Subscription, error)

	Delete(userID string, id string) error
}

type PaymentMethodDataSource interface {
	Create(
		userID string,
		label string,
		last4 string,
		brand model.PaymentMethodBrand,
		expMonth int,
		expYear int,
	) (*model.PaymentMethod, error)

	FindByID(userID string, id string) (*model.PaymentMethodWithSubscriptions, error)
	FindByUserID(userID string) ([]*model.PaymentMethodWithSubscriptions, error)

	Delete(userID string, id string) error
}

type UserDataSource interface {
	Create(email string) (*model.User, error)

	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)

	Delete(id string) error
}

type LoginTokenDataSource interface {
	Create(email string, hashedToken string) (*model.LoginToken, error)

	FindUnexpired(id string) (*model.LoginToken, error)

	Delete(id string) error
}

// --- Misc ---

type Hasher interface {
	Hash(input string) (string, error)
	Compare(input, hashed string) error
}

type Emailer interface {
	Send(from, to, subject, body string) error
}
