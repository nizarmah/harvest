package membership

import (
	"testing"
	"time"

	"github.com/whatis277/harvest/bean/internal/driver/postgres/postgrestest"
	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"
)

var (
	newMember = "00000000-0000-0000-0000-000000000001"

	activeMember  = "00000000-0000-0000-0003-000000000001"
	expiredMember = "00000000-0000-0000-0003-000000000002"
	nonMember     = "00000000-0000-0000-0003-000000000003"
)

func TestDataSource(t *testing.T) {
	db := postgrestest.DBTest(t)
	ds := New(db)

	t.Run("create", func(t *testing.T) {
		create(t, ds)
	})

	t.Run("find", func(t *testing.T) {
		find(t, ds)
	})

	t.Run("update", func(t *testing.T) {
		update(t, ds)
	})

	t.Run("delete", func(t *testing.T) {
		delete(t, ds)
	})
}

func create(t *testing.T, ds interfaces.MembershipDataSource) {
	t.Run("new_membership", func(t *testing.T) {
		createdAt := time.Now().Add(time.Hour).Truncate(time.Millisecond)

		membership, err := ds.Create(newMember, createdAt)
		if err != nil {
			t.Fatalf("failed to create membership: %s", err)
		}

		if membership.UserID != newMember {
			t.Errorf("expected user ID %s, got: %s", newMember, membership.UserID)
		}

		if !membership.CreatedAt.Equal(createdAt) {
			t.Errorf("expected created at %s, got: %s", createdAt, membership.CreatedAt)
		}

		if membership.ExpiresAt != nil {
			t.Errorf("expected nil expires at, got: %s", membership.ExpiresAt)
		}

		if err = ds.Delete(membership.UserID); err != nil {
			t.Fatalf("failed to cleanup membership: %s", err)
		}
	})

	t.Run("existing_membership", func(t *testing.T) {
		createdAt := time.Now().Add(time.Hour).Truncate(time.Millisecond)
		expiresAt := createdAt.Add(time.Hour).Truncate(time.Millisecond)

		membership, err := ds.Create(newMember, createdAt)
		if err != nil {
			t.Fatalf("failed to create membership: %s", err)
		}

		membership, err = ds.Update(membership.UserID, expiresAt)
		if err != nil {
			t.Fatalf("failed to update membership: %s", err)
		}

		newCreatedAt := createdAt.Add(time.Minute).Truncate(time.Millisecond)
		membership, err = ds.Create(membership.UserID, newCreatedAt)
		if err != nil {
			t.Fatalf("failed to create membership: %s", err)
		}

		if membership.UserID != newMember {
			t.Errorf("expected user ID %s, got: %s", newMember, membership.UserID)
		}

		if !membership.CreatedAt.Equal(newCreatedAt) {
			t.Errorf("expected created at %s, got: %s", newCreatedAt, membership.CreatedAt)
		}

		if membership.ExpiresAt != nil {
			t.Errorf("expected nil expires at, got: %s", membership.ExpiresAt)
		}

		if err = ds.Delete(membership.UserID); err != nil {
			t.Fatalf("failed to cleanup membership: %s", err)
		}
	})
}

func find(t *testing.T, ds interfaces.MembershipDataSource) {
	t.Run("existing_membership", func(t *testing.T) {
		membership, err := ds.Find(activeMember)
		if err != nil {
			t.Fatalf("failed to find membership: %s", err)
		}

		if membership.UserID != activeMember {
			t.Errorf("expected user ID %s, got: %s", activeMember, membership.UserID)
		}

		if membership.CreatedAt.IsZero() {
			t.Errorf("expected non-zero created at, got: %s", membership.CreatedAt)
		}

		if membership.ExpiresAt != nil {
			t.Errorf("expected nil expires at, got: %s", membership.ExpiresAt)
		}
	})

	t.Run("expired_membership", func(t *testing.T) {
		membership, err := ds.Find(expiredMember)
		if err != nil {
			t.Fatalf("failed to find membership: %s", err)
		}

		if membership.UserID != expiredMember {
			t.Errorf("expected user ID %s, got: %s", expiredMember, membership.UserID)
		}

		if membership.CreatedAt.IsZero() {
			t.Errorf("expected non-zero created at, got: %s", membership.CreatedAt)
		}

		if membership.ExpiresAt == nil {
			t.Errorf("expected non-nil expires at, got nil")
		}
	})

	t.Run("missing_membership", func(t *testing.T) {
		membership, err := ds.Find(nonMember)
		if err != nil {
			t.Fatalf("failed to find membership: %s", err)
		}

		if membership != nil {
			t.Errorf("expected nil membership, got: %v", membership)
		}
	})
}

func update(t *testing.T, ds interfaces.MembershipDataSource) {
	t.Run("existing_membership", func(t *testing.T) {
		createdAt := time.Now().Add(time.Hour).Truncate(time.Millisecond)
		expiresAt := createdAt.Add(time.Hour).Truncate(time.Millisecond)

		membership, err := ds.Create(newMember, createdAt)
		if err != nil {
			t.Fatalf("failed to create membership: %s", err)
		}

		membership, err = ds.Update(membership.UserID, expiresAt)
		if err != nil {
			t.Fatalf("failed to update membership: %s", err)
		}

		if membership.UserID != newMember {
			t.Errorf("expected user ID %s, got: %s", newMember, membership.UserID)
		}

		if !membership.CreatedAt.Equal(createdAt) {
			t.Errorf("expected created at %s, got: %s", createdAt, membership.CreatedAt)
		}

		if !membership.ExpiresAt.Equal(expiresAt) {
			t.Errorf("expected expires at %s, got: %s", expiresAt, membership.ExpiresAt)
		}

		if err = ds.Delete(membership.UserID); err != nil {
			t.Fatalf("failed to cleanup membership: %s", err)
		}
	})

	t.Run("missing_membership", func(t *testing.T) {
		expiresAt := time.Now().Add(time.Hour)

		membership, err := ds.Update(nonMember, expiresAt)
		if err != nil {
			t.Fatalf("failed to update membership: %s", err)
		}

		if membership != nil {
			t.Errorf("expected nil membership, got: %v", membership)
		}
	})
}

func delete(t *testing.T, ds interfaces.MembershipDataSource) {
	t.Run("existing_membership", func(t *testing.T) {
		membership, err := ds.Create(newMember, time.Now())
		if err != nil {
			t.Fatalf("failed to create membership: %s", err)
		}

		if err = ds.Delete(membership.UserID); err != nil {
			t.Fatalf("failed to delete membership: %s", err)
		}

		membership, err = ds.Find(membership.UserID)
		if err != nil {
			t.Fatalf("failed to find membership: %s", err)
		}

		if membership != nil {
			t.Errorf("expected nil membership, got: %v", membership)
		}
	})

	t.Run("missing_membership", func(t *testing.T) {
		err := ds.Delete(nonMember)
		if err != nil {
			t.Fatalf("failed to delete membership: %s", err)
		}
	})
}
