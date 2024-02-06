package paymentmethod

import (
	"context"
	"fmt"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase/interfaces"

	"harvest/bean/internal/driver/postgres"

	"github.com/jackc/pgx/v5"
)

type dataSource struct {
	db *postgres.DB
}

func New(db *postgres.DB) interfaces.PaymentMethodDataSource {
	return &dataSource{db}
}

func (ds *dataSource) Create(
	userID string,
	label string,
	last4 string,
	brand entity.PaymentMethodBrand,
	expMonth int,
	expYear int,
) (*entity.PaymentMethod, error) {
	method := &entity.PaymentMethod{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			("INSERT INTO payment_methods (user_id, label, last4, brand, exp_month, exp_year)"+
				" VALUES ($1, $2, $3, $4, $5, $6)"+
				" RETURNING *"),
			userID, label, last4, brand, expMonth, expYear,
		).
		Scan(
			&method.ID, &method.UserID,
			&method.Label, &method.Last4, &method.Brand, &method.ExpMonth, &method.ExpYear,
			&method.CreatedAt, &method.UpdatedAt,
		)

	if err != nil {
		return nil, fmt.Errorf("failed to create payment method: %w", err)
	}

	return method, nil
}

func (ds *dataSource) FindByID(userID string, id string) (*entity.PaymentMethod, error) {
	method := &entity.PaymentMethod{}

	err := ds.db.Pool.
		QueryRow(
			context.Background(),
			"SELECT * FROM payment_methods WHERE user_id = $1 AND id = $2",
			userID, id,
		).
		Scan(
			&method.ID, &method.UserID,
			&method.Label, &method.Last4, &method.Brand, &method.ExpMonth, &method.ExpYear,
			&method.CreatedAt, &method.UpdatedAt,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find payment method: %w", err)
	}

	return method, nil
}

func (ds *dataSource) FindByUserID(userID string) ([]*entity.PaymentMethod, error) {
	rows, err := ds.db.Pool.
		Query(
			context.Background(),
			"SELECT * FROM payment_methods WHERE user_id = $1",
			userID,
		)

	if err != nil {
		return nil, fmt.Errorf("failed to find payment methods: %w", err)
	}

	defer rows.Close()

	methods := []*entity.PaymentMethod{}
	for rows.Next() {
		method := &entity.PaymentMethod{}

		err := rows.Scan(
			&method.ID, &method.UserID,
			&method.Label, &method.Last4, &method.Brand, &method.ExpMonth, &method.ExpYear,
			&method.CreatedAt, &method.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan payment method: %w", err)
		}

		methods = append(methods, method)
	}

	return methods, nil
}

func (ds *dataSource) Delete(userID string, id string) error {
	_, err := ds.db.Pool.
		Exec(
			context.Background(),
			"DELETE FROM payment_methods WHERE user_id = $1 AND id = $2",
			userID, id,
		)

	if err != nil {
		return fmt.Errorf("failed to delete payment method: %w", err)
	}

	return nil
}
