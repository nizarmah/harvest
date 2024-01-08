package entity

import (
	"time"
)

type Subscription struct {
	ID int

	UserID          int
	PaymentMethodID int

	Amount int

	FreqVal  int
	FreqUnit string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentMethod struct{}

type User struct {
	ID int

	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginToken struct {
	ID int

	Email       string
	HashedToken []byte

	CreatedAt time.Time
	ExpiresAt time.Time
}
