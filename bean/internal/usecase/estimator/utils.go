package userdash

import (
	"harvest/bean/internal/entity/model"
)

func getSubscriptionEstimates(sub *model.Subscription) *model.Estimates {
	e := &model.Estimates{}

	switch sub.Period {
	case model.SubscriptionPeriodDaily:
		e = getEstimatesFromDaily(sub.Amount, sub.Interval)
	case model.SubscriptionPeriodWeekly:
		e = getEstimatesFromWeekly(sub.Amount, sub.Interval)
	case model.SubscriptionPeriodMonthly:
		e = getEstimatesFromMonthly(sub.Amount, sub.Interval)
	case model.SubscriptionPeriodYearly:
		e = getEstimatesFromYearly(sub.Amount, sub.Interval)
	}

	return e
}

func getEstimatesFromDaily(amount int, interval int) *model.Estimates {
	return &model.Estimates{
		Daily:   amount / interval,
		Weekly:  (amount * 7) / interval,
		Monthly: (amount * 30) / interval,
		Yearly:  (amount * 365) / interval,
	}
}

func getEstimatesFromWeekly(amount int, interval int) *model.Estimates {
	return &model.Estimates{
		Daily:   amount / (7 * interval),
		Weekly:  amount / interval,
		Monthly: (amount * 4) / interval,
		Yearly:  (amount * 52) / interval,
	}
}

func getEstimatesFromMonthly(amount int, interval int) *model.Estimates {
	return &model.Estimates{
		Daily:   amount / (30 * interval),
		Weekly:  amount / (4 * interval),
		Monthly: amount / interval,
		Yearly:  (amount * 12) / interval,
	}
}

func getEstimatesFromYearly(amount int, interval int) *model.Estimates {
	return &model.Estimates{
		Daily:   amount / (365 * interval),
		Weekly:  amount / (52 * interval),
		Monthly: amount / (12 * interval),
		Yearly:  amount / interval,
	}
}
