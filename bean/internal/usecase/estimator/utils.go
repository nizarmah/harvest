package userdash

import (
	"harvest/bean/internal/entity/model"
)

func getSubscriptionEstimates(sub *model.Subscription) *model.Estimates {
	e := &model.Estimates{}

	switch sub.Period {
	case model.SubscriptionPeriodDaily:
		e = getEstimatesFromDaily(sub.Amount)
	case model.SubscriptionPeriodWeekly:
		e = getEstimatesFromWeekly(sub.Amount)
	case model.SubscriptionPeriodMonthly:
		e = getEstimatesFromMonthly(sub.Amount)
	case model.SubscriptionPeriodYearly:
		e = getEstimatesFromYearly(sub.Amount)
	}

	return e
}

func getEstimatesFromDaily(amount int) *model.Estimates {
	return &model.Estimates{
		Daily:   amount,
		Weekly:  amount * 7,
		Monthly: amount * 30,
		Yearly:  amount * 365,
	}
}

func getEstimatesFromWeekly(amount int) *model.Estimates {
	return &model.Estimates{
		Daily:   amount / 7,
		Weekly:  amount,
		Monthly: amount * 4,
		Yearly:  amount * 52,
	}
}

func getEstimatesFromMonthly(amount int) *model.Estimates {
	return &model.Estimates{
		Daily:   amount / 30,
		Weekly:  amount / 4,
		Monthly: amount,
		Yearly:  amount * 12,
	}
}

func getEstimatesFromYearly(amount int) *model.Estimates {
	return &model.Estimates{
		Daily:   amount / 365,
		Weekly:  amount / 52,
		Monthly: amount / 12,
		Yearly:  amount,
	}
}
