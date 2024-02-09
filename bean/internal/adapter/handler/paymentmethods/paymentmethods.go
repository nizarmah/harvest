package landing

import (
	"fmt"
	"net/http"
	"net/url"

	"harvest/bean/internal/entity/model"
	"harvest/bean/internal/entity/viewmodel"

	estimatorUC "harvest/bean/internal/usecase/estimator"
	"harvest/bean/internal/usecase/paymentmethod"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	usecase   paymentmethod.UseCase
	estimator estimatorUC.UseCase

	view interfaces.PaymentMethodsView
}

func New(
	uc paymentmethod.UseCase,
	es estimatorUC.UseCase,
	view interfaces.PaymentMethodsView,
) http.Handler {
	return &handler{
		usecase:   uc,
		estimator: es,
		view:      view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methods, err := h.usecase.List("10000000-0000-0000-0000-000000000001")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	err = h.view.Render(w, h.makeViewData(methods))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}

func (h *handler) makeViewData(paymentMethods []*model.PaymentMethodWithSubscriptions) *viewmodel.PaymentMethodsViewData {
	methods := make([]viewmodel.PaymentMethodViewData, 0, len(paymentMethods))
	monthly, yearly := 0, 0

	for _, method := range paymentMethods {
		estimates := h.estimator.GetEstimates(method.Subscriptions)

		monthly += estimates.Monthly
		yearly += estimates.Yearly

		methods = append(methods, makeMethodViewData(method, estimates))
	}

	return &viewmodel.PaymentMethodsViewData{
		PaymentMethods:  methods,
		MonthlyEstimate: toDollarsString(monthly),
		YearlyEstimate:  toDollarsString(yearly),
	}
}

func makeMethodViewData(
	pm *model.PaymentMethodWithSubscriptions,
	estimates *model.Estimates,
) viewmodel.PaymentMethodViewData {
	method, subs := pm.PaymentMethod, pm.Subscriptions

	label := method.Label
	if label == "" {
		label = fmt.Sprintf("Card %s", method.Last4)
	}

	return viewmodel.PaymentMethodViewData{
		ID:       method.ID,
		Label:    label,
		Last4:    method.Last4,
		Brand:    string(method.Brand),
		ExpMonth: method.ExpMonth,
		ExpYear:  method.ExpYear,

		MonthlyEstimate: toDollarsString(estimates.Monthly),
		YearlyEstimate:  toDollarsString(estimates.Yearly),

		Subscriptions: makeSubsViewData(subs),
	}
}

func makeSubsViewData(subscriptions []*model.Subscription) []viewmodel.SubscriptionViewData {
	subs := make([]viewmodel.SubscriptionViewData, 0, len(subscriptions))
	for _, sub := range subscriptions {
		subs = append(subs, makeSubViewData(sub))
	}

	return subs
}

func makeSubViewData(subscription *model.Subscription) viewmodel.SubscriptionViewData {
	label := subscription.Label
	provider := subscription.Provider

	u, _ := url.Parse(provider)
	if label == "" && u != nil {
		label = u.Host
	}

	return viewmodel.SubscriptionViewData{
		ID:        subscription.ID,
		Label:     label,
		Provider:  provider,
		Amount:    toDollarsString(subscription.Amount),
		Frequency: toFrequencyString(subscription.Interval, subscription.Period),
	}
}

func toFrequencyString(interval int, period model.SubscriptionPeriod) string {
	if interval == 1 {
		return fmt.Sprintf("Every %s", period)
	}

	return fmt.Sprintf("Every %d %ss", interval, period)
}

func toDollarsString(amount int) string {
	dollars, cents := toDollars(amount)
	return fmt.Sprintf("$%d.%02d", dollars, cents)
}

func toDollars(amount int) (int, int) {
	return amount / 100, amount % 100
}
