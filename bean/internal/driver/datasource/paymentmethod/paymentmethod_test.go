package paymentmethod

import (
	"testing"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/postgres/postgrestest"
)

var (
	userWithMethodsId   = "00000000-0000-0000-0001-000000000001"
	userWithNoMethodsId = "00000000-0000-0000-0001-000000000002"

	methodId = "00000000-0000-0000-0000-000000000001"
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

func create(t *testing.T, ds usecase.PaymentMethodDataSource) {
	method, err := ds.Create(userWithMethodsId, "action-create", "1234", "brand", 12, 2024)
	if err != nil {
		t.Fatalf("failed to create payment method: %s", err)
	}

	if method.ID == "" {
		t.Error("expected payment method ID, got empty string")
	}

	if method.UserID != userWithMethodsId {
		t.Errorf("expected user ID: %s, got: %s", userWithMethodsId, method.UserID)
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

	if err := ds.Delete(method.ID); err != nil {
		t.Fatalf("failed to cleanup payment method: %s", err)
	}
}

func findById(t *testing.T, ds usecase.PaymentMethodDataSource) {
	t.Run("existing_payment_method", func(t *testing.T) {
		if _, err := ds.FindById(methodId); err != nil {
			t.Fatalf("failed to find payment method by id: %s", err)
		}
	})

	t.Run("nonexistent_payment_method", func(t *testing.T) {
		if _, err := ds.FindById("nonexistent"); err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func findByUserId(t *testing.T, ds usecase.PaymentMethodDataSource) {
	t.Run("has_payment_methods", func(t *testing.T) {
		methods, err := ds.FindByUserId(userWithMethodsId)
		if err != nil {
			t.Fatalf("failed to find payment methods: %s", err)
		}

		if len(methods) != 2 {
			t.Errorf("expected %d payment methods, got: %d", 2, len(methods))
		}

		for _, method := range methods {
			if method.UserID != userWithMethodsId {
				t.Errorf("expected user ID: %s, got: %s", userWithMethodsId, method.UserID)
			}
		}
	})

	t.Run("no_payment_methods", func(t *testing.T) {
		methods, err := ds.FindByUserId(userWithNoMethodsId)
		if err != nil {
			t.Fatalf("failed to find payment methods: %s", err)
		}

		if len(methods) != 0 {
			t.Errorf("expected %d payment methods, got: %d", 0, len(methods))
		}
	})
}

func delete(t *testing.T, ds usecase.PaymentMethodDataSource) {
	t.Run("existing_payment_method", func(t *testing.T) {
		method, err := ds.Create(userWithMethodsId, "action-delete", "1234", "brand", 12, 2024)
		if err != nil {
			t.Fatalf("failed to create payment method: %s", err)
		}

		if err = ds.Delete(method.ID); err != nil {
			t.Fatalf("failed to delete payment method: %s", err)
		}

		if _, err = ds.FindById(method.ID); err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("nonexistent_payment_method", func(t *testing.T) {
		if err := ds.Delete("nonexistent"); err == nil {
			t.Error("expected error, got nil")
		}
	})
}
