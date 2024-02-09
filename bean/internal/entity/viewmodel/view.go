package viewmodel

// --- View Data ---

type ViewData interface{}

// --- View Data --- Landing ---

type LandingViewData struct{}

// --- View Data --- Login ---

type LoginViewData struct {
	Email string
}

// --- View Data --- Payment Methods ---

type PaymentMethodsViewData struct {
	PaymentMethods []PaymentMethodViewData

	MonthlyEstimate string
	YearlyEstimate  string
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

	Subscriptions []SubscriptionViewData
}

type SubscriptionViewData struct {
	ID string

	Label     string
	Provider  string
	Amount    string
	Frequency string
}
