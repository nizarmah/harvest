package interfaces

import (
	"net/http"

	"harvest/bean/internal/entity/viewmodel"
)

// --- Views ---

type View[T any] interface {
	Render(http.ResponseWriter, *T) error
}

type LandingView View[viewmodel.LandingViewData]
type LoginView View[viewmodel.LoginViewData]
type HomeView View[viewmodel.HomeViewData]

type CreatePaymentMethodView View[viewmodel.CreatePaymentMethodViewData]
type DeletePaymentMethodView View[viewmodel.DeletePaymentMethodViewData]

type CreateSubscriptionView View[viewmodel.CreateSubscriptionViewData]
type DeleteSubscriptionView View[viewmodel.DeleteSubscriptionViewData]
