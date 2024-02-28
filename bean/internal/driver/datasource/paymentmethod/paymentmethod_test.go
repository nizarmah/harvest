package paymentmethod

import (
	"context"
	"testing"
	"time"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/whatis277/harvest/bean/internal/driver/postgres/postgrestest"
)

var (
	userWithMethodsID   = "00000000-0000-0000-0001-000000000001"
	userWithNoMethodsID = "00000000-0000-0000-0001-000000000002"

	methodID  = "00000000-0000-0000-0000-000000000001"
	missingID = "11111111-1111-1111-1111-111111111111"
)

func TestDataSouce(t *testing.T) {
	db := postgrestest.DBTest(t)
	ds := New(db)

	t.Run("create", func(t *testing.T) {
		create(t, ds)
	})

	t.Run("find_by_id", func(t *testing.T) {
		findByID(t, ds)
	})

	t.Run("find_by_user_id", func(t *testing.T) {
		findByUserID(t, ds)
	})

	t.Run("delete", func(t *testing.T) {
		delete(t, ds)
	})
}

func create(t *testing.T, ds interfaces.PaymentMethodDataSource) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	method, err := ds.Create(ctx, userWithMethodsID, "action-create", "1234", model.PaymentMethodBrandAmex, 12, 2024)
	if err != nil {
		t.Fatalf("failed to create payment method: %s", err)
	}

	if method.ID == "" {
		t.Error("expected payment method ID, got empty string")
	}

	if method.UserID != userWithMethodsID {
		t.Errorf("expected user ID: %s, got: %s", userWithMethodsID, method.UserID)
	}

	if method.Label != "action-create" {
		t.Errorf("expected label: %s, got: %s", "action-create", method.Label)
	}

	if method.Last4 != "1234" {
		t.Errorf("expected last4: %s, got: %s", "1234", method.Last4)
	}

	if method.Brand != model.PaymentMethodBrandAmex {
		t.Errorf("expected brand: %s, got: %s", model.PaymentMethodBrandAmex, method.Brand)
	}

	if method.ExpMonth != 12 {
		t.Errorf("expected expMonth: %d, got: %d", 12, method.ExpMonth)
	}

	if method.ExpYear != 2024 {
		t.Errorf("expected expYear: %d, got: %d", 2024, method.ExpYear)
	}

	if err := ds.Delete(ctx, userWithMethodsID, method.ID); err != nil {
		t.Fatalf("failed to cleanup payment method: %s", err)
	}
}

func findByID(t *testing.T, ds interfaces.PaymentMethodDataSource) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Run("existing_payment_method", func(t *testing.T) {
		method, err := ds.FindByID(ctx, userWithMethodsID, methodID)
		if err != nil {
			t.Fatalf("failed to find payment method by id: %s", err)
		}

		if method == nil {
			t.Fatalf("expected payment method, got nil")
		}

		if len(method.Subscriptions) != 1 {
			t.Errorf("expected %d subscriptions, got: %d", 1, len(method.Subscriptions))
		}

		for _, sub := range method.Subscriptions {
			if sub.ID == "" {
				t.Error("expected subscription ID, got empty string")
			}
		}
	})

	t.Run("not_own_payment_method", func(t *testing.T) {
		method, err := ds.FindByID(ctx, userWithNoMethodsID, methodID)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if method != nil {
			t.Errorf("expected nil payment method, got: %v", method)
		}
	})

	t.Run("missing_payment_method", func(t *testing.T) {
		method, err := ds.FindByID(ctx, userWithMethodsID, missingID)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if method != nil {
			t.Errorf("expected nil payment method, got: %v", method)
		}
	})
}

func findByUserID(t *testing.T, ds interfaces.PaymentMethodDataSource) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Run("has_payment_methods", func(t *testing.T) {
		methods, err := ds.FindByUserID(ctx, userWithMethodsID)
		if err != nil {
			t.Fatalf("failed to find payment methods: %s", err)
		}

		if len(methods) != 2 {
			t.Errorf("expected %d payment methods, got: %d", 2, len(methods))
		}

		for _, method := range methods {
			if method.PaymentMethod.UserID != userWithMethodsID {
				t.Errorf("expected user ID: %s, got: %s", userWithMethodsID, method.PaymentMethod.UserID)
			}

			if len(method.Subscriptions) != 1 {
				t.Errorf("expected %d subscriptions, got: %d", 1, len(method.Subscriptions))
			}

			for _, sub := range method.Subscriptions {
				if sub.ID == "" {
					t.Error("expected subscription ID, got empty string")
				}
			}
		}
	})

	t.Run("no_payment_methods", func(t *testing.T) {
		methods, err := ds.FindByUserID(ctx, userWithNoMethodsID)
		if err != nil {
			t.Fatalf("failed to find payment methods: %s", err)
		}

		if len(methods) != 0 {
			t.Errorf("expected %d payment methods, got: %d", 0, len(methods))
		}
	})
}

func delete(t *testing.T, ds interfaces.PaymentMethodDataSource) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Run("existing_payment_method", func(t *testing.T) {
		method, err := ds.Create(ctx, userWithMethodsID, "action-delete", "1234", model.PaymentMethodBrandAmex, 12, 2024)
		if err != nil {
			t.Fatalf("failed to create payment method: %s", err)
		}

		if err = ds.Delete(ctx, userWithMethodsID, method.ID); err != nil {
			t.Fatalf("failed to delete payment method: %s", err)
		}

		methodWithSubs, err := ds.FindByID(ctx, userWithMethodsID, method.ID)
		if err != nil {
			t.Fatalf("failed to find payment method: %s", err)
		}

		if methodWithSubs != nil {
			t.Errorf("expected nil payment method, got: %v", method)
		}
	})

	t.Run("not_own_payment_method", func(t *testing.T) {
		if err := ds.Delete(ctx, userWithNoMethodsID, methodID); err != nil {
			t.Fatalf("failed to delete payment method: %s", err)
		}

		method, err := ds.FindByID(ctx, userWithMethodsID, methodID)
		if err != nil {
			t.Fatalf("failed to find payment method: %s", err)
		}

		if method == nil {
			t.Errorf("expected payment method, got nil")
		}
	})

	t.Run("missing_payment_method", func(t *testing.T) {
		if err := ds.Delete(ctx, userWithMethodsID, missingID); err != nil {
			t.Fatalf("failed to delete payment method: %s", err)
		}
	})
}
