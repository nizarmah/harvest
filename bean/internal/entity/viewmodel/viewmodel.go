package viewmodel

// --- View Data ---

type ViewData interface{}

// --- View Data --- Landing ---

type LandingViewData struct{}

// --- View Data --- Login ---

type LoginViewData struct {
	Email string
}

// --- View Data --- Home ---

type HomeViewData struct {
	PaymentMethods []PaymentMethod

	MonthlyEstimate string
	YearlyEstimate  string
}

type PaymentMethod struct {
	ID string

	Label    string
	Last4    string
	Brand    string
	ExpMonth int
	ExpYear  int

	MonthlyEstimate string
	YearlyEstimate  string

	Subscriptions []Subscription
}

type Subscription struct {
	ID string

	Label     string
	Provider  string
	Amount    string
	Frequency string
}

// --- View Data --- Payment Method ---

type CreatePaymentMethodViewData struct{}
