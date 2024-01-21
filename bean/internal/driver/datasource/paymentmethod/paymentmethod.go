package paymentmethod

import (
	"errors"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/database"
)

type dataSource struct {
	db *database.DB
}

func New(db *database.DB) usecase.PaymentMethodDataSource {
	return &dataSource{db}
}

func (ds *dataSource) Create(input *entity.PaymentMethod) (*entity.PaymentMethod, error) {
	res, err := ds.db.Pool.Exec(
		"INSERT INTO payment_methods (user_id, label, last4, brand, exp_month, exp_year, is_default) VALUES (?, ?, ?, ?, ?, ?, ?)",
		input.UserID,
		input.Label,
		input.Last4,
		input.Brand,
		input.ExpMonth,
		input.ExpYear,
		input.IsDefault,
	)
	if err != nil {
		return nil, errors.New("error inserting payment method")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.New("error getting last insert id")
	}

	method := &entity.PaymentMethod{}

	err = ds.db.Pool.
		QueryRow("SELECT * FROM payment_methods WHERE id = ?", id).
		Scan(
			&method.ID, &method.UserID,
			&method.Label, &method.Last4, &method.Brand, &method.ExpMonth, &method.ExpYear,
			&method.IsDefault,
			&method.CreatedAt, &method.UpdatedAt,
		)
	if err != nil {
		return nil, errors.New("error getting inserted payment method")
	}

	return method, nil
}

func (ds *dataSource) FindByUserId(userId string) ([]*entity.PaymentMethod, error) {
	rows, err := ds.db.Pool.Query("SELECT * FROM payment_methods WHERE user_id = ?", userId)
	if err != nil {
		return nil, errors.New("error getting payment methods")
	}
	defer rows.Close()

	methods := []*entity.PaymentMethod{}

	for rows.Next() {
		method := &entity.PaymentMethod{}

		err := rows.Scan(
			&method.ID, &method.UserID,
			&method.Label, &method.Last4, &method.Brand, &method.ExpMonth, &method.ExpYear,
			&method.IsDefault,
			&method.CreatedAt, &method.UpdatedAt,
		)

		if err != nil {
			return nil, errors.New("error scanning payment method")
		}

		methods = append(methods, method)
	}

	return methods, nil
}

func (ds *dataSource) Delete(paymentMethod *entity.PaymentMethod) error {
	_, err := ds.db.Pool.Exec("DELETE FROM payment_methods WHERE id = ?", paymentMethod.ID)
	if err != nil {
		return errors.New("error deleting payment method")
	}

	return nil
}
