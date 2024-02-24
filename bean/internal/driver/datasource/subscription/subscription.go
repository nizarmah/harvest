package subscription

import (
	"context"
	"fmt"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/whatis277/harvest/bean/internal/driver/postgres"
)

type dataSource struct {
	db *postgres.DB
}

func New(db *postgres.DB) interfaces.SubscriptionDataSource {
	return &dataSource{
		db: db,
	}
}

func (ds *dataSource) Create(
	userID string,
	paymentMethodID string,
	label string,
	provider string,
	amount int,
	interval int,
	period model.SubscriptionPeriod,
) (*model.Subscription, error) {
	sub := &model.Subscription{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			("INSERT INTO subscriptions"+
				" (user_id, payment_method_id, label, provider, amount, interval, period)"+
				" VALUES ($1, $2, $3, $4, $5, $6, $7)"+
				" RETURNING *"),
			userID, paymentMethodID, label, provider, amount, interval, period,
		).
		Scan(
			&sub.ID, &sub.UserID, &sub.PaymentMethodID,
			&sub.Label, &sub.Provider,
			&sub.Amount, &sub.Interval, &sub.Period,
			&sub.CreatedAt, &sub.UpdatedAt,
		)

	if err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}

	return sub, nil
}

func (ds *dataSource) FindByID(userID string, id string) (*model.Subscription, error) {
	sub := &model.Subscription{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			("SELECT * FROM subscriptions"+
				" WHERE"+
				" subscriptions.user_id = $1"+
				" AND subscriptions.id = $2"),
			userID, id,
		).
		Scan(
			&sub.ID, &sub.UserID, &sub.PaymentMethodID,
			&sub.Label, &sub.Provider,
			&sub.Amount, &sub.Interval, &sub.Period,
			&sub.CreatedAt, &sub.UpdatedAt,
		)

	if err == postgres.ErrNowRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find subscription: %w", err)
	}

	return sub, nil
}

func (ds *dataSource) Delete(userID string, id string) error {
	_, err := ds.db.Pool.
		Exec(
			context.Background(),
			"DELETE FROM subscriptions WHERE user_id = $1 AND id = $2",
			userID, id,
		)

	if err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}

	return nil
}
