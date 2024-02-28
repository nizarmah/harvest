package interfaces

import (
	"context"
	"time"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

// --- Data Sources ---

type SubscriptionDataSource interface {
	Create(
		ctx context.Context,
		userID string,
		paymentMethodID string,
		label string,
		provider string,
		amount int,
		interval int,
		period model.SubscriptionPeriod,
	) (*model.Subscription, error)

	FindByID(
		ctx context.Context,
		userID string,
		id string,
	) (*model.Subscription, error)

	Delete(
		ctx context.Context,
		userID string,
		id string,
	) error
}

type PaymentMethodDataSource interface {
	Create(
		ctx context.Context,
		userID string,
		label string,
		last4 string,
		brand model.PaymentMethodBrand,
		expMonth int,
		expYear int,
	) (*model.PaymentMethod, error)

	FindByID(
		ctx context.Context,
		userID string,
		id string,
	) (*model.PaymentMethodWithSubscriptions, error)
	FindByUserID(
		ctx context.Context,
		userID string,
	) ([]*model.PaymentMethodWithSubscriptions, error)

	Delete(
		ctx context.Context,
		userID string,
		id string,
	) error
}

type UserDataSource interface {
	Create(
		ctx context.Context,
		email string,
	) (*model.User, error)

	FindById(
		ctx context.Context,
		id string,
	) (*model.User, error)
	FindByEmail(
		ctx context.Context,
		email string,
	) (*model.User, error)

	Delete(
		ctx context.Context,
		id string,
	) error
}

type LoginTokenDataSource interface {
	Create(
		ctx context.Context,
		email string,
		hashedToken string,
	) (*model.LoginToken, error)

	FindUnexpiredByEmail(
		ctx context.Context,
		email string,
	) (*model.LoginToken, error)
	FindUnexpiredByID(
		ctx context.Context,
		id string,
	) (*model.LoginToken, error)

	Delete(
		ctx context.Context,
		id string,
	) error
}

type MembershipDataSource interface {
	Create(
		ctx context.Context,
		userID string,
		createdAt time.Time,
	) (*model.Membership, error)

	Find(
		ctx context.Context,
		userID string,
	) (*model.Membership, error)

	Update(
		ctx context.Context,
		userID string,
		expiresAt time.Time,
	) (*model.Membership, error)

	Delete(
		ctx context.Context,
		userID string,
	) error
}

type SessionDataSource interface {
	Create(
		ctx context.Context,
		userID string,
		hashedToken string,
		duration time.Duration,
	) (*model.Session, error)

	FindByID(
		ctx context.Context,
		id string,
	) (*model.Session, error)

	Refresh(
		ctx context.Context,
		session *model.Session,
		duration time.Duration,
	) error

	Delete(
		ctx context.Context,
		id string,
	) error
}

// --- Misc ---

type Hasher interface {
	Hash(input string) (string, error)
	Compare(input, hashed string) error
}

type Emailer interface {
	Send(from, to, subject, body string) error
}
