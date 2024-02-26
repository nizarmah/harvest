package token

import (
	"testing"
	"time"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/whatis277/harvest/bean/internal/driver/postgres/postgrestest"
)

var (
	expiredID    = "00000000-0000-0000-0000-000000000001"
	expiredEmail = "expired"

	missingID    = "11111111-1111-1111-1111-111111111111"
	missingEmail = "missing"
)

func TestDataSource(t *testing.T) {
	db := postgrestest.DBTest(t)
	ds := New(db)

	t.Run("create", func(t *testing.T) {
		create(t, ds)
	})

	t.Run("find_unexpired_by_email", func(t *testing.T) {
		findUnexpiredByEmail(t, ds)
	})

	t.Run("find_unexpired_by_id", func(t *testing.T) {
		findUnexpiredByID(t, ds)
	})

	t.Run("delete", func(t *testing.T) {
		delete(t, ds)
	})
}

func create(t *testing.T, ds interfaces.LoginTokenDataSource) {
	t.Run("new_token", func(t *testing.T) {
		token, err := ds.Create("action-create", "hashed-token")
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

func findUnexpiredByEmail(t *testing.T, ds interfaces.LoginTokenDataSource) {
	t.Run("valid_token", func(t *testing.T) {
		token, err := ds.Create("action-find", "hashed-token")
		if err != nil {
			t.Fatalf("failed to create token: %s", err)
		}

		token, err = ds.FindUnexpiredByEmail(token.Email)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if token == nil {
			t.Fatalf("expected token, got nil")
		}

		if err = ds.Delete(token.ID); err != nil {
			t.Fatalf("failed to cleanup token: %s", err)
		}
	})

	t.Run("expired_token", func(t *testing.T) {
		token, err := ds.FindUnexpiredByEmail(expiredEmail)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if token != nil && token.ExpiresAt.Before(time.Now()) {
			t.Errorf("expected nil token, got: %v", token)
		}
	})

	t.Run("missing_token", func(t *testing.T) {
		token, err := ds.FindUnexpiredByEmail(missingEmail)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if token != nil {
			t.Errorf("expected nil token, got: %v", token)
		}
	})
}

func findUnexpiredByID(t *testing.T, ds interfaces.LoginTokenDataSource) {
	t.Run("valid_token", func(t *testing.T) {
		token, err := ds.Create("action-find", "hashed-token")
		if err != nil {
			t.Fatalf("failed to create token: %s", err)
		}

		token, err = ds.FindUnexpiredByID(token.ID)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if token == nil {
			t.Fatalf("expected token, got nil")
		}

		if err = ds.Delete(token.ID); err != nil {
			t.Fatalf("failed to cleanup token: %s", err)
		}
	})

	t.Run("expired_token", func(t *testing.T) {
		token, err := ds.FindUnexpiredByID(expiredID)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if token != nil && token.ExpiresAt.Before(time.Now()) {
			t.Errorf("expected nil token, got: %v", token)
		}
	})

	t.Run("missing_token", func(t *testing.T) {
		token, err := ds.FindUnexpiredByID(missingID)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if token != nil {
			t.Errorf("expected nil token, got: %v", token)
		}
	})
}

func delete(t *testing.T, ds interfaces.LoginTokenDataSource) {
	t.Run("existing_token", func(t *testing.T) {
		token, err := ds.Create("action-delete", "hashed-token")
		if err != nil {
			t.Fatalf("failed to create token: %s", err)
		}

		if err = ds.Delete(token.ID); err != nil {
			t.Fatalf("failed to delete token: %s", err)
		}

		token, err = ds.FindUnexpiredByID(token.ID)
		if err != nil {
			t.Fatalf("failed to find token: %s", err)
		}

		if token != nil {
			t.Errorf("expected nil token, got: %v", token)
		}
	})

	t.Run("missing_token", func(t *testing.T) {
		if err := ds.Delete(missingID); err != nil {
			t.Fatalf("failed to delete token: %s", err)
		}
	})
}
