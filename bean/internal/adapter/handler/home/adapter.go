package home

import (
	"fmt"
	"net/url"

	"harvest/bean/internal/entity/model"
	"harvest/bean/internal/entity/viewmodel"
)

func toPaymentMethodViewModel(
	pm *model.PaymentMethodWithSubscriptions,
	estimates *model.Estimates,
) viewmodel.PaymentMethod {
	method, subs := pm.PaymentMethod, pm.Subscriptions

	label := method.Label
	if label == "" {
		label = fmt.Sprintf("Card %s", method.Last4)
	}

	return viewmodel.PaymentMethod{
		ID:       method.ID,
		Label:    label,
		Last4:    method.Last4,
		Brand:    string(method.Brand),
		ExpMonth: method.ExpMonth,
		ExpYear:  method.ExpYear,

		MonthlyEstimate: toDollarsString(estimates.Monthly),
		YearlyEstimate:  toDollarsString(estimates.Yearly),

		Subscriptions: toSubscriptionsViewModel(subs),
	}
}

func toSubscriptionsViewModel(subscriptions []*model.Subscription) []viewmodel.Subscription {
	subs := make([]viewmodel.Subscription, 0, len(subscriptions))
	for _, sub := range subscriptions {
		subs = append(subs, toSubscriptionViewModel(sub))
	}

	return subs
}

func toSubscriptionViewModel(subscription *model.Subscription) viewmodel.Subscription {
	label := subscription.Label
	provider := subscription.Provider

	u, _ := url.Parse(provider)
	if label == "" && u != nil {
		label = u.Host
	}

	return viewmodel.Subscription{
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
