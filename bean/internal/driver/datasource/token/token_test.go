package token

import (
	"testing"
	"time"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/postgres/postgrestest"
)

func TestDataSource(t *testing.T) {
	db := postgrestest.DBTest(t)
	ds := New(db)

	t.Run("create", func(t *testing.T) {
		create(t, ds)
	})

	t.Run("find_unexpired", func(t *testing.T) {
		findUnexpired(t, ds)
	})

	t.Run("delete", func(t *testing.T) {
		delete(t, ds)
	})
}

func create(t *testing.T, ds usecase.LoginTokenDataSource) {
	t.Run("new_token", func(t *testing.T) {
		token, err := ds.Create("action-new", "hashed-token")
		if err != nil {
			t.Fatalf("failed to create token: %s", err)
		}

		if expiry := token.ExpiresAt.Sub(token.CreatedAt).Minutes(); expiry != 10 {
			t.Errorf("expected 10 minute expiry, got: %v", expiry)
		}

		if err = ds.Delete(token.ID); err != nil {
			t.Fatalf("failed to cleanup token: %s", err)
		}
	})

	t.Run("existing_token", func(t *testing.T) {
		old, err := ds.Create("action-overwrite", "old-hashed-token")
		if err != nil {
			t.Fatalf("failed to create token: %s", err)
		}

		new, err := ds.Create("action-overwrite", "new-hashed-token")
		if err != nil {
			t.Fatalf("failed to create token: %s", err)
		}

		if old.ID != new.ID {
			t.Errorf("expected same token ID, got: %s", new.ID)
		}

		if old.CreatedAt == new.CreatedAt {
			t.Errorf("expected new creation time, got: %v", old.CreatedAt)
		}

		if old.ExpiresAt == new.ExpiresAt {
			t.Errorf("expected new expiry time, got: %v", old.ExpiresAt)
		}

		if old.HashedToken == new.HashedToken {
			t.Errorf("expected new hashed token, got: %s", old.HashedToken)
		}

		if err = ds.Delete(new.ID); err != nil {
			t.Fatalf("failed to cleanup token: %s", err)
		}
	})
}

func findUnexpired(t *testing.T, ds usecase.LoginTokenDataSource) {
	t.Run("valid_token", func(t *testing.T) {
		token, err := ds.Create("action-find", "hashed-token")
		if err != nil {
			t.Fatalf("failed to create token: %s", err)
		}

		if _, err = ds.FindUnexpired(token.ID); err != nil {
			t.Fatalf("failed to find token: %s", err)
		}

		if err = ds.Delete(token.ID); err != nil {
			t.Fatalf("failed to cleanup token: %s", err)
		}
	})

	t.Run("expired_token", func(t *testing.T) {
		token, _ := ds.FindUnexpired("00000000-0000-0000-0000-000000000001")
		if token != nil && token.ExpiresAt.Before(time.Now()) {
			t.Errorf("expected nil token, got: %v", token)
		}
	})

	t.Run("nonexistent_token", func(t *testing.T) {
		if _, err := ds.FindUnexpired("nonexistent"); err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func delete(t *testing.T, ds usecase.LoginTokenDataSource) {
	t.Run("existing_token", func(t *testing.T) {
		token, err := ds.Create("action-delete", "hashed-token")
		if err != nil {
			t.Fatalf("failed to create token: %s", err)
		}

		if err = ds.Delete(token.ID); err != nil {
			t.Fatalf("failed to delete token: %s", err)
		}

		if _, err = ds.FindUnexpired(token.ID); err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("nonexistent_token", func(t *testing.T) {
		if err := ds.Delete("nonexistent"); err == nil {
			t.Error("expected error, got nil")
		}
	})
}
