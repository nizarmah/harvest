package userdash

import (
	"harvest/bean/internal/entity"
)

func getSubscriptionEstimates(sub *entity.Subscription) *entity.Estimates {
	e := &entity.Estimates{}

	switch sub.Period {
	case entity.SubscriptionPeriodDaily:
		e = getEstimatesFromDaily(sub.Amount)
	case entity.SubscriptionPeriodWeekly:
		e = getEstimatesFromWeekly(sub.Amount)
	case entity.SubscriptionPeriodMonthly:
		e = getEstimatesFromMonthly(sub.Amount)
	case entity.SubscriptionPeriodYearly:
		e = getEstimatesFromYearly(sub.Amount)
	}

	return e
}

func getEstimatesFromDaily(amount int) *entity.Estimates {
	return &entity.Estimates{
		Daily:   amount,
		Weekly:  amount * 7,
		Monthly: amount * 30,
		Yearly:  amount * 365,
	}
}

func getEstimatesFromWeekly(amount int) *entity.Estimates {
	return &entity.Estimates{
		Daily:   amount / 7,
		Weekly:  amount,
		Monthly: amount * 4,
		Yearly:  amount * 52,
	}
}

func getEstimatesFromMonthly(amount int) *entity.Estimates {
	return &entity.Estimates{
		Daily:   amount / 30,
		Weekly:  amount / 4,
		Monthly: amount,
		Yearly:  amount * 12,
	}
}

func getEstimatesFromYearly(amount int) *entity.Estimates {
	return &entity.Estimates{
		Daily:   amount / 365,
		Weekly:  amount / 52,
		Monthly: amount / 12,
		Yearly:  amount,
	}
}
