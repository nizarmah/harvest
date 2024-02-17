package subscription

import (
	"fmt"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

func validateLabel(label string) error {
	if len(label) > 255 {
		return fmt.Errorf("label must be less than 255 chars")
	}

	return nil
}

func validateProvider(provider string) error {
	if len(provider) > 255 {
		return fmt.Errorf("provider must be less than 255 chars")
	}

	return nil
}

func validateAmount(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	return nil
}

func validateInterval(interval int) error {
	if interval <= 0 {
		return fmt.Errorf("interval must be greater than 0")
	}

	if interval > 365 {
		return fmt.Errorf("interval must be less than 366")
	}

	return nil
}

func validatePeriod(period model.SubscriptionPeriod) error {
	switch period {
	case model.SubscriptionPeriodDaily,
		model.SubscriptionPeriodWeekly,
		model.SubscriptionPeriodMonthly,
		model.SubscriptionPeriodYearly:
		return nil
	default:
		return fmt.Errorf(
			"period must be: %s, %s, %s, %s",
			model.SubscriptionPeriodDaily,
			model.SubscriptionPeriodWeekly,
			model.SubscriptionPeriodMonthly,
			model.SubscriptionPeriodYearly,
		)
	}
}
