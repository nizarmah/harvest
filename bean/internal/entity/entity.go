package entity

import (
	"time"
)

type Subscription struct {
	ID     string
	UserID string

	Amount        int
	Frequency     int
	PaymentMethod string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID string

	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Token struct {
	ID     string
	UserID string

	AccessToken string

	CreatedAt time.Time
	ExpiresAt time.Time
}
