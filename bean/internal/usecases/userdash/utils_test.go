package userdash

import (
	"testing"
)

func TestGetEstimatesFromDaily(t *testing.T) {
	amount := 100
	estimates := getEstimatesFromDaily(amount)

	if estimates.Daily != amount {
		t.Errorf("daily estimate should be %d, got %d", amount, estimates.Daily)
	}

	if estimates.Weekly != amount*7 {
		t.Errorf("weekly estimate should be %d, got %d", amount*7, estimates.Weekly)
	}

	if estimates.Monthly != amount*30 {
		t.Errorf("monthly estimate should be %d, got %d", amount*30, estimates.Monthly)
	}

	if estimates.Yearly != amount*365 {
		t.Errorf("yearly estimate should be %d, got %d", amount*365, estimates.Yearly)
	}
}

func TestGetEstimatesFromWeekly(t *testing.T) {
	amount := 700
	estimates := getEstimatesFromWeekly(amount)

	if estimates.Daily != amount/7 {
		t.Errorf("daily estimate should be %d, got %d", amount/7, estimates.Daily)
	}

	if estimates.Weekly != amount {
		t.Errorf("weekly estimate should be %d, got %d", amount, estimates.Weekly)
	}

	if estimates.Monthly != amount*4 {
		t.Errorf("monthly estimate should be %d, got %d", amount*4, estimates.Monthly)
	}

	if estimates.Yearly != amount*52 {
		t.Errorf("yearly estimate should be %d, got %d", amount*52, estimates.Yearly)
	}
}

func TestGetEstimatesFromMonthly(t *testing.T) {
	amount := 2500
	estimates := getEstimatesFromMonthly(amount)

	if estimates.Daily != amount/30 {
		t.Errorf("daily estimate should be %d, got %d", amount/30, estimates.Daily)
	}

	if estimates.Weekly != amount/4 {
		t.Errorf("weekly estimate should be %d, got %d", amount/4, estimates.Weekly)
	}

	if estimates.Monthly != amount {
		t.Errorf("monthly estimate should be %d, got %d", amount, estimates.Monthly)
	}

	if estimates.Yearly != amount*12 {
		t.Errorf("yearly estimate should be %d, got %d", amount*12, estimates.Yearly)
	}
}

func TestGetEstimatesFromYearly(t *testing.T) {
	amount := 100000
	estimates := getEstimatesFromYearly(amount)

	t.Log(estimates)

	if estimates.Daily != amount/365 {
		t.Errorf("daily estimate should be %d, got %d", amount/365, estimates.Daily)
	}

	if estimates.Weekly != amount/52 {
		t.Errorf("weekly estimate should be %d, got %d", amount/52, estimates.Weekly)
	}

	if estimates.Monthly != amount/12 {
		t.Errorf("monthly estimate should be %d, got %d", amount/12, estimates.Monthly)
	}

	if estimates.Yearly != amount {
		t.Errorf("yearly estimate should be %d, got %d", amount, estimates.Yearly)
	}
}
