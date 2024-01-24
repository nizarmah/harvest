package userdash

import (
	"fmt"

	"harvest/bean/internal/entity"
	"harvest/bean/internal/usecases/interfaces"
)

type UseCase struct {
	subscriptions interfaces.SubscriptionDataSource
}

func (u *UseCase) GetMonthlyApproxTotal(userID string) (int, error) {
	subscriptions, err := u.subscriptions.FindByUserID(userID)
	if err != nil {
		return -1, fmt.Errorf("failed to get subscriptions: %w", err)
	}

	amount := 0
	for _, subscription := range subscriptions {
		switch subscription.Period {
		case entity.SubscriptionPeriodDaily:
			amount += subscription.Amount * 30
		case entity.SubscriptionPeriodWeekly:
			amount += subscription.Amount * 4
		case entity.SubscriptionPeriodMonthly:
			amount += subscription.Amount
		case entity.SubscriptionPeriodYearly:
			amount += subscription.Amount / 12
		}
	}

	return amount, nil
}

func (u *UseCase) GetYearlyApproxTotal(userID string) (int, error) {
	subscriptions, err := u.subscriptions.FindByUserID(userID)
	if err != nil {
		return -1, fmt.Errorf("failed to get subscriptions: %w", err)
	}

	amount := 0
	for _, subscription := range subscriptions {
		switch subscription.Period {
		case entity.SubscriptionPeriodDaily:
			amount += subscription.Amount * 365
		case entity.SubscriptionPeriodWeekly:
			amount += subscription.Amount * 52
		case entity.SubscriptionPeriodMonthly:
			amount += subscription.Amount * 12
		case entity.SubscriptionPeriodYearly:
			amount += subscription.Amount
		}
	}

	return amount, nil
}
