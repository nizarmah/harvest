package userdash

import (
	"testing"
)

func TestGetEstimatesFromDaily(t *testing.T) {
	amount := 100
	estimates := getEstimatesFromDaily(amount, 2)

	if estimates.Daily != amount/2 {
		t.Errorf("daily estimate should be %d, got %d", amount/2, estimates.Daily)
	}

	if estimates.Weekly != (amount*7)/2 {
		t.Errorf("weekly estimate should be %d, got %d", (amount*7)/2, estimates.Weekly)
	}

	if estimates.Monthly != (amount*30)/2 {
		t.Errorf("monthly estimate should be %d, got %d", (amount*30)/2, estimates.Monthly)
	}

	if estimates.Yearly != (amount*365)/2 {
		t.Errorf("yearly estimate should be %d, got %d", (amount*365)/2, estimates.Yearly)
	}
}

func TestGetEstimatesFromWeekly(t *testing.T) {
	amount := 700
	estimates := getEstimatesFromWeekly(amount, 7)

	if estimates.Daily != amount/49 {
		t.Errorf("daily estimate should be %d, got %d", amount/49, estimates.Daily)
	}

	if estimates.Weekly != amount/7 {
		t.Errorf("weekly estimate should be %d, got %d", amount/7, estimates.Weekly)
	}

	if estimates.Monthly != amount*4/7 {
		t.Errorf("monthly estimate should be %d, got %d", amount*4/7, estimates.Monthly)
	}

	if estimates.Yearly != amount*52/7 {
		t.Errorf("yearly estimate should be %d, got %d", amount*52/7, estimates.Yearly)
	}
}

func TestGetEstimatesFromMonthly(t *testing.T) {
	amount := 2500
	estimates := getEstimatesFromMonthly(amount, 5)

	if estimates.Daily != amount/(30*5) {
		t.Errorf("daily estimate should be %d, got %d", amount/(30*5), estimates.Daily)
	}

	if estimates.Weekly != amount/(4*5) {
		t.Errorf("weekly estimate should be %d, got %d", amount/(4*5), estimates.Weekly)
	}

	if estimates.Monthly != amount/5 {
		t.Errorf("monthly estimate should be %d, got %d", amount/5, estimates.Monthly)
	}

	if estimates.Yearly != amount*12/5 {
		t.Errorf("yearly estimate should be %d, got %d", amount*12/5, estimates.Yearly)
	}
}

func TestGetEstimatesFromYearly(t *testing.T) {
	amount := 100000
	estimates := getEstimatesFromYearly(amount, 2)

	if estimates.Daily != amount/(365*2) {
		t.Errorf("daily estimate should be %d, got %d", amount/(365*2), estimates.Daily)
	}

	if estimates.Weekly != amount/(52*2) {
		t.Errorf("weekly estimate should be %d, got %d", amount/(52*2), estimates.Weekly)
	}

	if estimates.Monthly != amount/(12*2) {
		t.Errorf("monthly estimate should be %d, got %d", amount/(12*2), estimates.Monthly)
	}

	if estimates.Yearly != amount/2 {
		t.Errorf("yearly estimate should be %d, got %d", amount/2, estimates.Yearly)
	}
}
