package model

import (
	"encoding/json"
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

// --- Database --- Membership ---

type Membership struct {
	UserID string

	CreatedAt time.Time
	ExpiresAt *time.Time
}

// --- Cache ---

// --- Cache --- Session ---

type Session struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`

	HashedToken string `json:"hashed_token"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (s *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

// --- Cookie ---

// --- Cookie --- Session Token ---

type SessionToken struct {
	ID string `json:"id"`

	Token string `json:"token"`

	ExpiresAt time.Time `json:"expires_at"`
}

func (s *SessionToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SessionToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

// --- Error ---

// --- Error --- User Input ---

type UserInputError error

// --- Misc ---

type Estimates struct {
	Daily   int
	Weekly  int
	Monthly int
	Yearly  int
}
