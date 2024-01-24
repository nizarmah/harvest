package subscription

import (
	"fmt"
	"harvest/bean/internal/entity"
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

	return nil
}

func validatePeriod(period entity.SubscriptionPeriod) error {
	switch period {
	case entity.SubscriptionPeriodDaily,
		entity.SubscriptionPeriodWeekly,
		entity.SubscriptionPeriodMonthly,
		entity.SubscriptionPeriodYearly:
		return nil
	default:
		return fmt.Errorf("period must be day, week, month or year")
	}
}
