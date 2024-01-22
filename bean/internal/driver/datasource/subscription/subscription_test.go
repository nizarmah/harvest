package subscription

import (
	"testing"

	"harvest/bean/internal/usecases/interfaces"

	"harvest/bean/internal/driver/postgres/postgrestest"
)

var (
	userWithSubsId   = "00000000-0000-0000-0002-000000000001"
	userWithNoSubsId = "00000000-0000-0000-0002-000000000002"

	methodId = "00000000-0000-0000-0001-000000000001"

	subId = "00000000-0000-0000-0000-000000000001"
)

func TestDataSouce(t *testing.T) {
	db := postgrestest.DBTest(t)
	ds := New(db)

	t.Run("create", func(t *testing.T) {
		create(t, ds)
	})

	t.Run("find_by_id", func(t *testing.T) {
		findById(t, ds)
	})

	t.Run("find_by_user_id", func(t *testing.T) {
		findByUserId(t, ds)
	})

	t.Run("delete", func(t *testing.T) {
		delete(t, ds)
	})
}

func create(t *testing.T, ds interfaces.SubscriptionDataSource) {
	sub, err := ds.Create(userWithSubsId, methodId, "action-create", "bean", 1000, 1, "month")
	if err != nil {
		t.Fatalf("failed to create subscription: %s", err)
	}

	if sub.ID == "" {
		t.Error("expected subscription ID, got empty string")
	}

	if sub.UserID != userWithSubsId {
		t.Errorf("expected user ID: %s, got: %s", userWithSubsId, sub.UserID)
	}

	if sub.PaymentMethodID != methodId {
		t.Errorf("expected payment method ID: %s, got: %s", methodId, sub.PaymentMethodID)
	}

	if sub.Label != "action-create" {
		t.Errorf("expected label: %s, got: %s", "action-create", sub.Label)
	}

	if sub.Provider != "bean" {
		t.Errorf("expected provider: %s, got: %s", "bean", sub.Provider)
	}

	if sub.Amount != 1000 {
		t.Errorf("expected amount: %d, got: %d", 1000, sub.Amount)
	}

	if sub.Interval != 1 {
		t.Errorf("expected interval: %d, got: %d", 1, sub.Interval)
	}

	if sub.Period != "month" {
		t.Errorf("expected period: %s, got: %s", "month", sub.Period)
	}

	if err := ds.Delete(sub.ID); err != nil {
		t.Fatalf("failed to cleanup subscription: %s", err)
	}
}

func findById(t *testing.T, ds interfaces.SubscriptionDataSource) {
	t.Run("existing_subscription", func(t *testing.T) {
		if _, err := ds.FindById(subId); err != nil {
			t.Fatalf("failed to find subscription by id: %s", err)
		}
	})

	t.Run("nonexistent_subscription", func(t *testing.T) {
		if _, err := ds.FindById("nonexistent"); err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func findByUserId(t *testing.T, ds interfaces.SubscriptionDataSource) {
	t.Run("has_subscriptions", func(t *testing.T) {
		subs, err := ds.FindByUserId(userWithSubsId)
		if err != nil {
			t.Fatalf("failed to find subscriptions by user id: %s", err)
		}

		if len(subs) != 2 {
			t.Errorf("expected %d subscriptions, got: %d", 2, len(subs))
		}

		for _, sub := range subs {
			if sub.UserID != userWithSubsId {
				t.Errorf("expected user ID: %s, got: %s", userWithSubsId, sub.UserID)
			}
		}
	})

	t.Run("no_subscriptions", func(t *testing.T) {
		subs, err := ds.FindByUserId(userWithNoSubsId)
		if err != nil {
			t.Fatalf("failed to find subscriptions by user id: %s", err)
		}

		if len(subs) != 0 {
			t.Errorf("expected %d subscriptions, got: %d", 0, len(subs))
		}
	})
}

func delete(t *testing.T, ds interfaces.SubscriptionDataSource) {
	t.Run("existing_subscription", func(t *testing.T) {
		sub, err := ds.Create(userWithSubsId, methodId, "action-delete", "bean", 1000, 1, "month")
		if err != nil {
			t.Fatalf("failed to create subscription: %s", err)
		}

		if err = ds.Delete(sub.ID); err != nil {
			t.Fatalf("failed to delete subscription: %s", err)
		}

		if _, err = ds.FindById(sub.ID); err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("nonexistent_subscription", func(t *testing.T) {
		if err := ds.Delete("nonexistent"); err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
