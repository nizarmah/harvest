package user

import (
	"testing"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/whatis277/harvest/bean/internal/driver/postgres/postgrestest"
)

var (
	userID    = "00000000-0000-0000-0000-000000000001"
	missingID = "11111111-1111-1111-1111-111111111111"
)

func TestDataSource(t *testing.T) {
	db := postgrestest.DBTest(t)
	ds := New(db)

	t.Run("create", func(t *testing.T) {
		create(t, ds)
	})

	t.Run("find_by_id", func(t *testing.T) {
		findById(t, ds)
	})

	t.Run("find_by_email", func(t *testing.T) {
		findByEmail(t, ds)
	})

	t.Run("delete", func(t *testing.T) {
		delete(t, ds)
	})
}

func create(t *testing.T, ds interfaces.UserDataSource) {
	t.Run("new_user", func(t *testing.T) {
		user, err := ds.Create("action-create")
		if err != nil {
			t.Fatalf("failed to create user: %s", err)
		}

		if user.ID == "" {
			t.Error("expected user ID, got empty string")
		}

		if user.Email != "action-create" {
			t.Errorf("expected email: %s, got: %s", "action-create", user.Email)
		}

		if user.CreatedAt.IsZero() {
			t.Error("expected creation time, got zero time")
		}

		if err = ds.Delete(user.ID); err != nil {
			t.Fatalf("failed to cleanup user: %s", err)
		}
	})

	t.Run("existing_user", func(t *testing.T) {
		user, err := ds.Create("action-reject")
		if err != nil {
			t.Fatalf("failed to create user: %s", err)
		}

		if _, err = ds.Create("action-reject"); err == nil {
			t.Errorf("expected error, got nil")
		}

		if err = ds.Delete(user.ID); err != nil {
			t.Fatalf("failed to cleanup user: %s", err)
		}
	})
}

func findById(t *testing.T, ds interfaces.UserDataSource) {
	t.Run("existing_user", func(t *testing.T) {
		user, err := ds.FindById(userID)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if user == nil {
			t.Errorf("expected user, got nil")
		}
	})

	t.Run("missing_user", func(t *testing.T) {
		user, err := ds.FindById(missingID)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if user != nil {
			t.Errorf("expected nil user, got: %v", user)
		}
	})
}

func findByEmail(t *testing.T, ds interfaces.UserDataSource) {
	t.Run("existing_user", func(t *testing.T) {
		user, err := ds.FindByEmail("user-1")
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if user == nil {
			t.Errorf("expected user, got nil")
		}
	})

	t.Run("missing_user", func(t *testing.T) {
		user, err := ds.FindByEmail("missing")
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if user != nil {
			t.Errorf("expected nil user, got: %v", user)
		}
	})
}

func delete(t *testing.T, ds interfaces.UserDataSource) {
	t.Run("existing_user", func(t *testing.T) {
		user, err := ds.Create("action-delete")
		if err != nil {
			t.Fatalf("failed to create user: %s", err)
		}

		if err = ds.Delete(user.ID); err != nil {
			t.Fatalf("failed to delete user: %s", err)
		}

		user, err = ds.FindById(user.ID)
		if err != nil {
			t.Fatalf("failed to find user: %s", err)
		}

		if user != nil {
			t.Errorf("expected nil user, got: %v", user)
		}
	})

	t.Run("missing_user", func(t *testing.T) {
		if err := ds.Delete(missingID); err != nil {
			t.Fatalf("failed to delete user: %s", err)
		}
	})
}
