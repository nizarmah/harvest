package membership

import (
	"context"
	"fmt"
	"time"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/whatis277/harvest/bean/internal/driver/postgres"
)

type dataSource struct {
	db *postgres.DB
}

func New(db *postgres.DB) interfaces.MembershipDataSource {
	return &dataSource{db}
}

func (ds *dataSource) Create(
	ctx context.Context,
	userID string,
	createdAt time.Time,
) (*model.Membership, error) {
	membership := &model.Membership{}

	err := ds.db.Pool.
		QueryRow(
			ctx,
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

func (ds *dataSource) Find(
	ctx context.Context,
	userID string,
) (*model.Membership, error) {
	membership := &model.Membership{}

	err := ds.db.Pool.
		QueryRow(
			ctx,
			("SELECT * FROM memberships"+
				" WHERE user_id = $1"),
			userID,
		).
		Scan(
			&membership.UserID,
			&membership.CreatedAt, &membership.ExpiresAt,
		)

	if err == postgres.ErrNowRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find membership: %w", err)
	}

	return membership, nil
}

func (ds *dataSource) Update(
	ctx context.Context,
	userID string,
	expiresAt time.Time,
) (*model.Membership, error) {
	membership := &model.Membership{}

	err := ds.db.Pool.
		QueryRow(
			ctx,
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

	if err == postgres.ErrNowRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to update membership: %w", err)
	}

	return membership, nil
}

func (ds *dataSource) Delete(
	ctx context.Context,
	userID string,
) error {
	_, err := ds.db.Pool.
		Exec(
			ctx,
			"DELETE FROM memberships WHERE user_id = $1",
			userID,
		)

	if err != nil {
		return fmt.Errorf("failed to delete membership: %w", err)
	}

	return nil
}
