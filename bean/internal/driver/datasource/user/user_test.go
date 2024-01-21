package user

import (
	"testing"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/postgres/postgrestest"
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

func create(t *testing.T, ds usecase.UserDataSource) {
	t.Run("new_user", func(t *testing.T) {
		user, err := ds.Create("action-new")
		if err != nil {
			t.Fatalf("failed to create user: %s", err)
		}

		if user.ID == "" {
			t.Error("expected user ID, got empty string")
		}

		if user.Email != "action-new" {
			t.Errorf("expected email 'action-new', got: %s", user.Email)
		}

		if user.CreatedAt.IsZero() {
			t.Error("expected creation time, got zero time")
		}

		err = ds.Delete(user.ID)
		if err != nil {
			t.Fatalf("failed to cleanup user: %s", err)
		}
	})

	t.Run("existing_user", func(t *testing.T) {
		user, err := ds.Create("action-reject")
		if err != nil {
			t.Fatalf("failed to create user: %s", err)
		}

		_, err = ds.Create("action-reject")
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		err = ds.Delete(user.ID)
		if err != nil {
			t.Fatalf("failed to cleanup user: %s", err)
		}
	})
}

func findById(t *testing.T, ds usecase.UserDataSource) {
	t.Run("existing_user", func(t *testing.T) {
		user, err := ds.Create("action-find-by-id")
		if err != nil {
			t.Fatalf("failed to create user: %s", err)
		}

		found, err := ds.FindById(user.ID)
		if err != nil {
			t.Fatalf("failed to find user by id: %s", err)
		}

		if found.ID != user.ID {
			t.Errorf("expected same token ID, got: %s", found.ID)
		}

		err = ds.Delete(user.ID)
		if err != nil {
			t.Fatalf("failed to cleanup user: %s", err)
		}
	})

	t.Run("nonexistent_user", func(t *testing.T) {
		_, err := ds.FindById("nonexistent")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func findByEmail(t *testing.T, ds usecase.UserDataSource) {
	t.Run("existing_user", func(t *testing.T) {
		user, err := ds.Create("action-find-by-email")
		if err != nil {
			t.Fatalf("failed to create user: %s", err)
		}

		found, err := ds.FindByEmail(user.Email)
		if err != nil {
			t.Fatalf("failed to find user by email: %s", err)
		}

		if found.ID != user.ID {
			t.Errorf("expected same token ID, got: %s", found.ID)
		}

		err = ds.Delete(user.ID)
		if err != nil {
			t.Fatalf("failed to cleanup user: %s", err)
		}
	})

	t.Run("nonexistent_user", func(t *testing.T) {
		_, err := ds.FindByEmail("nonexistent")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func delete(t *testing.T, ds usecase.UserDataSource) {
	t.Run("existing_user", func(t *testing.T) {
		user, err := ds.Create("action-delete")
		if err != nil {
			t.Fatalf("failed to create user: %s", err)
		}

		err = ds.Delete(user.ID)
		if err != nil {
			t.Fatalf("failed to delete user: %s", err)
		}

		_, err = ds.FindById(user.ID)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("nonexistent_user", func(t *testing.T) {
		err := ds.Delete("nonexistent")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
