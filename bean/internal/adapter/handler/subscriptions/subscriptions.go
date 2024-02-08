package landing

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase/subscription"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	usecase subscription.UseCase

	view interfaces.SubscriptionsView
}

func New(
	usecase subscription.UseCase,
	view interfaces.SubscriptionsView,
) http.Handler {
	return &handler{
		usecase: usecase,
		view:    view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	subs, err := h.usecase.List("10000000-0000-0000-0000-000000000001")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	err = h.view.Render(w, makeViewData(subs))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}

func makeViewData(subscriptions []*entity.SubscriptionWithPaymentMethod) *entity.SubscriptionsViewData {
	var subs = make([]entity.SubscriptionViewData, 0, len(subscriptions))
	for _, sub := range subscriptions {
		subs = append(subs, makeSubViewData(sub))
	}

	return &entity.SubscriptionsViewData{
		Subscriptions: subs,
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
