package entity

import (
	"time"
)

// --- Database ---

// --- Database --- Subscription ---

type SubscriptionWithPaymentMethod struct {
	Subscription  *Subscription
	PaymentMethod *PaymentMethod
}

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

// --- View Data ---

type ViewData interface{}

// --- View Data --- Landing ---

type LandingViewData struct{}

// --- View Data --- Login ---

type LoginViewData struct {
	Email string
}

// --- View Data --- Subscriptions ---

type SubscriptionsViewData struct {
	Subscriptions []SubscriptionViewData

	MonthlyEstimate string
	YearlyEstimate  string
}

type SubscriptionViewData struct {
	ID string

	Label     string
	Provider  string
	Amount    string
	Frequency string

	PaymentMethodID string

	PaymentMethodLabel    string
	PaymentMethodLast4    string
	PaymentMethodExpMonth int
	PaymentMethodExpYear  int
}

// --- View Data --- Payment Methods ---

type PaymentMethodsViewData struct {
	PaymentMethods []PaymentMethodViewData
}

type PaymentMethodViewData struct {
	ID string

	Label    string
	Last4    string
	Brand    string
	ExpMonth int
	ExpYear  int

	MonthlyEstimate string
	YearlyEstimate  string
}

// --- Misc ---

type Estimates struct {
	Daily   int
	Weekly  int
	Monthly int
	Yearly  int
}
