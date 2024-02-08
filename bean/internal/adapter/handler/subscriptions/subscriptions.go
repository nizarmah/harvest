package landing

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity"

	estimatorUC "harvest/bean/internal/usecase/estimator"
	"harvest/bean/internal/usecase/subscription"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	usecase   subscription.UseCase
	estimator estimatorUC.UseCase

	view interfaces.SubscriptionsView
}

func New(
	uc subscription.UseCase,
	es estimatorUC.UseCase,
	view interfaces.SubscriptionsView,
) http.Handler {
	return &handler{
		usecase:   uc,
		estimator: es,
		view:      view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	subs, err := h.usecase.List("10000000-0000-0000-0000-000000000001")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	err = h.view.Render(w, h.makeViewData(subs))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}

func (h *handler) makeViewData(subscriptions []*entity.SubscriptionWithPaymentMethod) *entity.SubscriptionsViewData {
	subs := make([]*entity.Subscription, 0, len(subscriptions))
	viewdata := make([]entity.SubscriptionViewData, 0, len(subscriptions))

	for _, sub := range subscriptions {
		subs = append(subs, sub.Subscription)
		viewdata = append(viewdata, makeSubViewData(sub))
	}

	estimates := h.estimator.GetEstimates(subs)

	monthlyDollars, monthlyCents := estimates.Monthly/100, estimates.Monthly%100
	yearlyDollars, yearlyCents := estimates.Yearly/100, estimates.Yearly%100

	return &entity.SubscriptionsViewData{
		Subscriptions: viewdata,

		MonthlyEstimate: fmt.Sprintf("$%d.%02d", monthlyDollars, monthlyCents),
		YearlyEstimate:  fmt.Sprintf("$%d.%02d", yearlyDollars, yearlyCents),
	}
}

func makeSubViewData(sub *entity.SubscriptionWithPaymentMethod) entity.SubscriptionViewData {
	subscription, method := sub.Subscription, sub.PaymentMethod

	dollars, cents := subscription.Amount/100, subscription.Amount%100

	frequency := makeFrequency(subscription.Interval, subscription.Period)

	return entity.SubscriptionViewData{
		ID:        subscription.ID,
		Label:     subscription.Label,
		Provider:  subscription.Provider,
		Amount:    fmt.Sprintf("$%d.%02d", dollars, cents),
		Frequency: frequency,

		PaymentMethodID:       method.ID,
		PaymentMethodLabel:    method.Label,
		PaymentMethodLast4:    method.Last4,
		PaymentMethodExpMonth: method.ExpMonth,
		PaymentMethodExpYear:  method.ExpYear,
	}
}

func makeFrequency(interval int, period entity.SubscriptionPeriod) string {
	if interval == 1 {
		return fmt.Sprintf("Every %s", period)
	}

	return fmt.Sprintf("Every %d %ss", interval, period)
}
