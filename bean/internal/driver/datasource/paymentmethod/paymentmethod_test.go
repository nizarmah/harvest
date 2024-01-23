package paymentmethod

import (
	"testing"

	"harvest/bean/internal/usecases/interfaces"

	"harvest/bean/internal/driver/postgres/postgrestest"
)

var (
	userWithMethodsID   = "00000000-0000-0000-0001-000000000001"
	userWithNoMethodsID = "00000000-0000-0000-0001-000000000002"

	methodID      = "00000000-0000-0000-0000-000000000001"
	nonexistentID = "11111111-1111-1111-1111-111111111111"
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
	method, err := ds.Create(userWithMethodsID, "action-create", "1234", "brand", 12, 2024)
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

	if method.Brand != "brand" {
		t.Errorf("expected brand: %s, got: %s", "brand", method.Brand)
	}

	if method.ExpMonth != 12 {
		t.Errorf("expected expMonth: %d, got: %d", 12, method.ExpMonth)
	}

	if method.ExpYear != 2024 {
		t.Errorf("expected expYear: %d, got: %d", 2024, method.ExpYear)
	}

	if err := ds.Delete(userWithMethodsID, method.ID); err != nil {
		t.Fatalf("failed to cleanup payment method: %s", err)
	}
}

func findByID(t *testing.T, ds interfaces.PaymentMethodDataSource) {
	t.Run("existing_payment_method", func(t *testing.T) {
		if _, err := ds.FindByID(userWithMethodsID, methodID); err != nil {
			t.Fatalf("failed to find payment method by id: %s", err)
		}
	})

	t.Run("not_own_payment_method", func(t *testing.T) {
		if _, err := ds.FindByID(userWithNoMethodsID, methodID); err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("nonexistent_payment_method", func(t *testing.T) {
		if _, err := ds.FindByID(userWithMethodsID, nonexistentID); err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func findByUserID(t *testing.T, ds interfaces.PaymentMethodDataSource) {
	t.Run("has_payment_methods", func(t *testing.T) {
		methods, err := ds.FindByUserID(userWithMethodsID)
		if err != nil {
			t.Fatalf("failed to find payment methods: %s", err)
		}

		if len(methods) != 2 {
			t.Errorf("expected %d payment methods, got: %d", 2, len(methods))
		}

		for _, method := range methods {
			if method.UserID != userWithMethodsID {
				t.Errorf("expected user ID: %s, got: %s", userWithMethodsID, method.UserID)
			}
		}
	})

	t.Run("no_payment_methods", func(t *testing.T) {
		methods, err := ds.FindByUserID(userWithNoMethodsID)
		if err != nil {
			t.Fatalf("failed to find payment methods: %s", err)
		}

		if len(methods) != 0 {
			t.Errorf("expected %d payment methods, got: %d", 0, len(methods))
		}
	})
}

func delete(t *testing.T, ds interfaces.PaymentMethodDataSource) {
	t.Run("existing_payment_method", func(t *testing.T) {
		method, err := ds.Create(userWithMethodsID, "action-delete", "1234", "brand", 12, 2024)
		if err != nil {
			t.Fatalf("failed to create payment method: %s", err)
		}

		if err = ds.Delete(userWithMethodsID, method.ID); err != nil {
			t.Fatalf("failed to delete payment method: %s", err)
		}

		if _, err = ds.FindByID(userWithMethodsID, method.ID); err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("not_own_payment_method", func(t *testing.T) {
		if err := ds.Delete(userWithNoMethodsID, methodID); err != nil {
			t.Fatalf("failed to delete payment method: %s", err)
		}

		if method, _ := ds.FindByID(userWithMethodsID, methodID); method == nil {
			t.Error("expected payment method, got nil")
		}
	})

	t.Run("nonexistent_payment_method", func(t *testing.T) {
		if err := ds.Delete(userWithMethodsID, nonexistentID); err != nil {
			t.Fatalf("failed to delete payment method: %s", err)
		}
	})
}
