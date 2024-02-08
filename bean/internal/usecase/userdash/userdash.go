package userdash

import (
	"harvest/bean/internal/entity"
)

type UseCase struct{}

func (u *UseCase) GetEstimates(subs []*entity.Subscription) *entity.Estimates {
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

	return estimates
}
