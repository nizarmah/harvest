package session

import (
	"context"
	"testing"
	"time"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/whatis277/harvest/bean/internal/driver/redis"
	"github.com/whatis277/harvest/bean/internal/driver/redis/redistest"
)

func TestDataSource(t *testing.T) {
	cache := redistest.CacheTest(t)
	ds := New(cache)

	t.Run("create", func(t *testing.T) {
		create(t, ds, cache)
	})

	t.Run("find_by_id", func(t *testing.T) {
		findById(t, ds)
	})

	t.Run("refresh", func(t *testing.T) {
		refresh(t, ds, cache)
	})

	t.Run("delete", func(t *testing.T) {
		delete(t, ds)
	})
}

func create(t *testing.T, ds interfaces.SessionDataSource, cache *redis.Cache) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	t.Run("new_session", func(t *testing.T) {
		session, err := ds.Create("action-create", "hashed-token", 10*time.Second)
		if err != nil {
			t.Fatalf("failed to create session: %s", err)
		}

		if session.ID == "" {
			t.Error("expected session ID, got empty string")
		}

		if session.UserID != "action-create" {
			t.Errorf("expected user ID: %s, got: %s", "action-create", session.UserID)
		}

		if session.HashedToken != "hashed-token" {
			t.Errorf("expected hashed token: %s, got: %s", "hashed-token", session.HashedToken)
		}

		if session.CreatedAt.IsZero() {
			t.Error("expected created at, got zero value")
		}

		if session.UpdatedAt.IsZero() {
			t.Error("expected updated at, got zero value")
		}

		if session.ExpiresAt.IsZero() {
			t.Error("expected expires at, got zero value")
		}

		duration := session.ExpiresAt.Sub(session.CreatedAt)
		if duration != 10*time.Second {
			t.Errorf("expected expires at to be 10 seconds after created at, got: %s", duration)
		}

		ttl := cache.Client.TTL(ctx, session.ID).Val()
		if ttl != 10*time.Second {
			t.Errorf("expected session to expire in: %s, got: %s", 10*time.Second, ttl)
		}

		if err = ds.Delete(session.ID); err != nil {
			t.Fatalf("failed to cleanup session: %s", err)
		}
	})
}

func findById(t *testing.T, ds interfaces.SessionDataSource) {
	t.Run("existing_session", func(t *testing.T) {
		session, err := ds.Create("action-find", "hashed-token", 10*time.Second)
		if err != nil {
			t.Fatalf("failed to create session: %s", err)
		}

		session, err = ds.FindByID(session.ID)
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if session == nil {
			t.Fatalf("expected session, got nil")
		}

		if session.ID == "" {
			t.Error("expected session ID, got empty string")
		}

		if session.UserID != "action-find" {
			t.Errorf("expected user ID: %s, got: %s", "action-find", session.UserID)
		}

		if session.HashedToken != "hashed-token" {
			t.Errorf("expected hashed token: %s, got: %s", "hashed-token", session.HashedToken)
		}

		if err = ds.Delete(session.ID); err != nil {
			t.Fatalf("failed to cleanup session: %s", err)
		}
	})

	t.Run("missing_session", func(t *testing.T) {
		session, err := ds.FindByID("missing-id")
		if err != nil {
			t.Fatalf("expected nil error, got: %s", err)
		}

		if session != nil {
			t.Errorf("expected nil session, got: %v", session)
		}
	})
}

func refresh(t *testing.T, ds interfaces.SessionDataSource, cache *redis.Cache) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	t.Run("existing_session", func(t *testing.T) {
		session, err := ds.Create("action-refresh", "hashed-token", 10*time.Second)
		if err != nil {
			t.Fatalf("failed to create session: %s", err)
		}

		duration := session.ExpiresAt.Sub(session.UpdatedAt)
		if duration != 10*time.Second {
			t.Errorf("expected expires at to be 10 seconds after updated at, got: %s", duration)
		}

		ttl := cache.Client.TTL(ctx, session.ID).Val()
		if ttl != 10*time.Second {
			t.Errorf("expected session to expire in: %s, got: %s", 10*time.Second, ttl)
		}

		err = ds.Refresh(session, 20*time.Second)
		if err != nil {
			t.Fatalf("failed to refresh session: %s", err)
		}

		duration = session.ExpiresAt.Sub(session.UpdatedAt)
		if duration != 20*time.Second {
			t.Errorf("expected expires at to be 20 seconds after updated at, got: %s", duration)
		}

		ttl = cache.Client.TTL(ctx, session.ID).Val()
		if ttl != 20*time.Second {
			t.Errorf("expected session to expire in: %s, got: %s", 20*time.Second, ttl)
		}

		if err = ds.Delete(session.ID); err != nil {
			t.Fatalf("failed to delete session: %s", err)
		}
	})
}

func delete(t *testing.T, ds interfaces.SessionDataSource) {
	t.Run("existing_session", func(t *testing.T) {
		session, err := ds.Create("action-delete", "hashed-token", 10*time.Second)
		if err != nil {
			t.Fatalf("failed to create session: %s", err)
		}

		if err = ds.Delete(session.ID); err != nil {
			t.Fatalf("failed to delete session: %s", err)
		}

		session, err = ds.FindByID(session.ID)
		if err != nil {
			t.Fatalf("failed to find session: %s", err)
		}

		if session != nil {
			t.Errorf("expected nil session, got: %v", session)
		}
	})

	t.Run("missing_session", func(t *testing.T) {
		if err := ds.Delete("missing-id"); err != nil {
			t.Fatalf("failed to delete session: %s", err)
		}
	})
}
