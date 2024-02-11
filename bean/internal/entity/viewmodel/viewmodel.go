package viewmodel

// --- View Models ---

// --- View Models --- Payment Method ---

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

// --- View Models --- Subscription ---

type Subscription struct {
	ID string

	Label     string
	Provider  string
	Amount    string
	Frequency string
}

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

// --- View Data --- Payment Method ---

type CreatePaymentMethodViewData struct {
	Error string

	Form CreatePaymentMethodForm
}

type DeletePaymentMethodViewData struct {
	PaymentMethod PaymentMethod
}

type CreatePaymentMethodForm struct {
	Label    string
	Last4    string
	Brand    string
	ExpMonth int
	ExpYear  int
}

// --- View Data --- Subscription ---

type CreateSubscriptionViewData struct {
	Error string

	Form CreateSubscriptionForm
}

type DeleteSubscriptionViewData struct {
	Subscription Subscription
}

type CreateSubscriptionForm struct {
	PaymentMethodID string

	Label    string
	Provider string
	Amount   float32
	Interval int
	Period   string
}
