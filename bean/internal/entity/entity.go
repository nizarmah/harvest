package entity

import (
	"time"
)

// --- Database ---

type SubscriptionPeriod string

const (
	SubscriptionPeriodDaily   SubscriptionPeriod = "day"
	SubscriptionPeriodWeekly  SubscriptionPeriod = "week"
	SubscriptionPeriodMonthly SubscriptionPeriod = "month"
	SubscriptionPeriodYearly  SubscriptionPeriod = "year"
)

type Subscription struct {
	ID              string
	UserID          string
	PaymentMethodID string

	Label    string
	Provider string
	Amount   int
	Interval int
	Period   SubscriptionPeriod

	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentMethod struct {
	ID     string
	UserID string

	Label    string
	Last4    string
	Brand    string
	ExpMonth int
	ExpYear  int

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

// --- Misc ---

type Estimates struct {
	Daily   int
	Weekly  int
	Monthly int
	Yearly  int
}
