package entity

import (
	"time"
)

type Subscription struct {
	ID int

	UserID          string
	PaymentMethodID int

	Amount int

	FreqVal  int
	FreqUnit string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentMethod struct {
	ID int

	UserID string

	Label    string
	Last4    string
	Brand    string
	ExpMonth int
	ExpYear  int

	IsDefault bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID string

	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginToken struct {
	ID string

	Email       string
	HashedToken string

	CreatedAt time.Time
	ExpiresAt time.Time
}
