package membership

import (
	"context"
	"fmt"
	"time"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/whatis277/harvest/bean/internal/driver/postgres"

	"github.com/jackc/pgx/v5"
)

type dataSource struct {
	db *postgres.DB
}

func New(db *postgres.DB) interfaces.MembershipDataSource {
	return &dataSource{db}
}

func (ds *dataSource) Create(userID string, createdAt time.Time) (*model.Membership, error) {
	membership := &model.Membership{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			("INSERT INTO memberships"+
				" (user_id, created_at)"+
				" VALUES ($1, $2)"+
				" ON CONFLICT (user_id) DO UPDATE"+
				" SET"+
				" created_at = $2,"+
				" expires_at = NULL"+
				" RETURNING *"),
			userID, createdAt,
		).
		Scan(
			&membership.UserID,
			&membership.CreatedAt, &membership.ExpiresAt,
		)

	if err != nil {
		return nil, fmt.Errorf("failed to create membership: %w", err)
	}

	return membership, nil
}

func (ds *dataSource) Find(userID string) (*model.Membership, error) {
	membership := &model.Membership{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			("SELECT * FROM memberships"+
				" WHERE user_id = $1"),
			userID,
		).
		Scan(
			&membership.UserID,
			&membership.CreatedAt, &membership.ExpiresAt,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find membership: %w", err)
	}

	return membership, nil
}

func (ds *dataSource) Update(userID string, expiresAt time.Time) (*model.Membership, error) {
	membership := &model.Membership{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			("UPDATE memberships"+
				" SET expires_at = $2"+
				" WHERE user_id = $1"+
				" RETURNING *"),
			userID, expiresAt,
		).
		Scan(
			&membership.UserID,
			&membership.CreatedAt, &membership.ExpiresAt,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to update membership: %w", err)
	}

	return membership, nil
}

func (ds *dataSource) Delete(userID string) error {
	_, err := ds.db.Pool.
		Exec(
			context.Background(),
			"DELETE FROM memberships WHERE user_id = $1",
			userID,
		)

	if err != nil {
		return fmt.Errorf("failed to delete membership: %w", err)
	}

	return nil
}
