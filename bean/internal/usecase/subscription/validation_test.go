package subscription

import (
	"strings"
	"testing"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

func TestValidateLabel(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		if err := validateLabel("label"); err != nil {
			t.Errorf("expected nil, got: %s", err)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		if err := validateLabel(
			strings.Repeat("a", 256),
		); err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestValidateProvider(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		if err := validateProvider("provider"); err != nil {
			t.Errorf("expected nil, got: %s", err)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		if err := validateProvider(
			strings.Repeat("a", 256),
		); err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestValidateAmount(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		if err := validateAmount(1); err != nil {
			t.Errorf("expected nil, got: %s", err)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		if err := validateAmount(0); err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestValidateInterval(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		if err := validateInterval(1); err != nil {
			t.Errorf("expected nil, got: %s", err)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		if err := validateInterval(0); err == nil {
			t.Error("expected error, got nil")
		}

		if err := validateInterval(366); err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestValidatePeriod(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tests := []model.SubscriptionPeriod{
			model.SubscriptionPeriodDaily,
			model.SubscriptionPeriodWeekly,
			model.SubscriptionPeriodMonthly,
			model.SubscriptionPeriodYearly,
		}

		for _, test := range tests {
			if err := validatePeriod(test); err != nil {
				t.Errorf("expected nil, got: %s", err)
			}
		}
	})

	t.Run("invalid", func(t *testing.T) {
		if err := validatePeriod("invalid"); err == nil {
			t.Error("expected error, got nil")
		}
	})
}
