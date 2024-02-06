package userdash

import (
	"fmt"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase/interfaces"
)

type UseCase struct {
	subscriptions interfaces.SubscriptionDataSource
}

func (u *UseCase) GetEstimates(userID string) (*entity.Estimates, error) {
	subs, err := u.subscriptions.FindByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscriptions: %w", err)
	}

	estimates := &entity.Estimates{
		Daily:   0,
		Weekly:  0,
		Monthly: 0,
		Yearly:  0,
	}

	for _, sub := range subs {
		e := getSubscriptionEstimates(sub)

		estimates.Daily += e.Daily
		estimates.Weekly += e.Weekly
		estimates.Monthly += e.Monthly
		estimates.Yearly += e.Yearly
	}

	return estimates, nil
}
