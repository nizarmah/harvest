package landing

import (
	"fmt"
	"net/http"
	"net/url"

	"harvest/bean/internal/entity"

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

func (h *handler) makeViewData(paymentMethods []*entity.PaymentMethodWithSubscriptions) *entity.PaymentMethodsViewData {
	viewdata := make([]entity.PaymentMethodViewData, 0, len(paymentMethods))
	monthly, yearly := 0, 0

	for _, method := range paymentMethods {
		viewdata = append(viewdata, h.makeMethodViewData(method))

		estimates := h.estimator.GetEstimates(method.Subscriptions)

		monthly += estimates.Monthly
		yearly += estimates.Yearly
	}

	monthlyDollars, monthlyCents := monthly/100, monthly%100
	yearlyDollars, yearlyCents := yearly/100, yearly%100

	return &entity.PaymentMethodsViewData{
		PaymentMethods:  viewdata,
		MonthlyEstimate: fmt.Sprintf("$%d.%02d", monthlyDollars, monthlyCents),
		YearlyEstimate:  fmt.Sprintf("$%d.%02d", yearlyDollars, yearlyCents),
	}
}

func (h *handler) makeMethodViewData(pm *entity.PaymentMethodWithSubscriptions) entity.PaymentMethodViewData {
	method, subs := pm.PaymentMethod, pm.Subscriptions

	estimates := h.estimator.GetEstimates(subs)

	monthlyDollars, monthlyCents := estimates.Monthly/100, estimates.Monthly%100
	yearlyDollars, yearlyCents := estimates.Yearly/100, estimates.Yearly%100

	label := method.Label
	if label == "" {
		label = fmt.Sprintf("Card %s", method.Last4)
	}

	return entity.PaymentMethodViewData{
		ID:       method.ID,
		Label:    label,
		Last4:    method.Last4,
		Brand:    string(method.Brand),
		ExpMonth: method.ExpMonth,
		ExpYear:  method.ExpYear,

		MonthlyEstimate: fmt.Sprintf("$%d.%02d", monthlyDollars, monthlyCents),
		YearlyEstimate:  fmt.Sprintf("$%d.%02d", yearlyDollars, yearlyCents),

		Subscriptions: h.makeSubsViewData(subs),
	}
}

func (h *handler) makeSubsViewData(subs []*entity.Subscription) []entity.SubscriptionViewData {
	viewdata := make([]entity.SubscriptionViewData, 0, len(subs))

	for _, sub := range subs {
		viewdata = append(viewdata, h.makeSubViewData(sub))
	}

	return viewdata
}

func (h *handler) makeSubViewData(subscription *entity.Subscription) entity.SubscriptionViewData {
	dollars, cents := subscription.Amount/100, subscription.Amount%100

	frequency := makeFrequency(subscription.Interval, subscription.Period)

	label := subscription.Label
	provider := subscription.Provider

	u, _ := url.Parse(provider)
	if label == "" && u != nil {
		label = u.Host
	}

	return entity.SubscriptionViewData{
		ID:        subscription.ID,
		Label:     label,
		Provider:  provider,
		Amount:    fmt.Sprintf("$%d.%02d", dollars, cents),
		Frequency: frequency,
	}
}

func makeFrequency(interval int, period entity.SubscriptionPeriod) string {
	if interval == 1 {
		return fmt.Sprintf("Every %s", period)
	}

	return fmt.Sprintf("Every %d %ss", interval, period)
}
