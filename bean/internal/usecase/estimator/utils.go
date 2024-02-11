package userdash

import (
	"harvest/bean/internal/entity/model"
)

func getSubscriptionEstimates(sub *model.Subscription) *model.Estimates {
	e := &model.Estimates{}

	amount := sub.Amount / sub.Interval

	switch sub.Period {
	case model.SubscriptionPeriodDaily:
		e = getEstimatesFromDaily(amount)
	case model.SubscriptionPeriodWeekly:
		e = getEstimatesFromWeekly(amount)
	case model.SubscriptionPeriodMonthly:
		e = getEstimatesFromMonthly(amount)
	case model.SubscriptionPeriodYearly:
		e = getEstimatesFromYearly(amount)
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
