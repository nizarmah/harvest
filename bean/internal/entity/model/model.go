package model

import (
	"fmt"
	"time"
)

// --- Database ---

// --- Database --- Subscription ---

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

type SubscriptionPeriod string

const (
	SubscriptionPeriodDaily   SubscriptionPeriod = "day"
	SubscriptionPeriodWeekly  SubscriptionPeriod = "week"
	SubscriptionPeriodMonthly SubscriptionPeriod = "month"
	SubscriptionPeriodYearly  SubscriptionPeriod = "year"
)

// --- Database --- Payment Method ---

type PaymentMethodWithSubscriptions struct {
	PaymentMethod *PaymentMethod
	Subscriptions []*Subscription
}

type PaymentMethod struct {
	ID     string
	UserID string

	Label    string
	Last4    string
	Brand    PaymentMethodBrand
	ExpMonth int
	ExpYear  int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentMethodBrand string

const (
	PaymentMethodBrandAmex       PaymentMethodBrand = "amex"
	PaymentMethodBrandMastercard PaymentMethodBrand = "mastercard"
	PaymentMethodBrandVisa       PaymentMethodBrand = "visa"
)

// --- Database --- User ---

type User struct {
	ID string

	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// --- Database --- Login Token ---

type LoginToken struct {
	ID string

	Email       string
	HashedToken string

	CreatedAt time.Time
	ExpiresAt time.Time
}

// --- Errors ---

type Error struct {
	Label   string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Label, e.Message)
}

// --- Misc ---

type Estimates struct {
	Daily   int
	Weekly  int
	Monthly int
	Yearly  int
}
