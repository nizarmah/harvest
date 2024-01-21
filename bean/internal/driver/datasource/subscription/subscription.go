package subscription

import (
	"errors"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/database"
)

type dataSource struct {
	db *database.DB
}

func New(db *database.DB) usecase.SubscriptionDataSource {
	return &dataSource{
		db: db,
	}
}

func (ds *dataSource) Create(input *entity.Subscription) (*entity.Subscription, error) {
	res, err := ds.db.Pool.Exec(
		"INSERT INTO subscriptions (user_id, payment_method_id, amount, freq_val, freq_unit) VALUES (?, ?, ?, ?, ?)",
		input.UserID,
		input.PaymentMethodID,
		input.Amount,
		input.FreqVal,
		input.FreqUnit,
	)
	if err != nil {
		return nil, errors.New("error inserting subscription")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.New("error getting last insert id")
	}

	sub := &entity.Subscription{}

	err = ds.db.Pool.
		QueryRow("SELECT * FROM subscriptions WHERE id = ?", id).
		Scan(
			&sub.ID, &sub.UserID, &sub.PaymentMethodID,
			&sub.Amount, &sub.FreqVal, &sub.FreqUnit,
			&sub.CreatedAt, &sub.UpdatedAt,
		)
	if err != nil {
		return nil, errors.New("error getting inserted subscription")
	}

	return sub, nil
}

func (ds *dataSource) FindByUserId(userId string) ([]*entity.Subscription, error) {
	subs := []*entity.Subscription{}

	rows, err := ds.db.Pool.Query("SELECT * FROM subscriptions WHERE user_id = ?", userId)
	if err != nil {
		return nil, errors.New("error getting subscriptions")
	}
	defer rows.Close()

	for rows.Next() {
		s := &entity.Subscription{}

		err = rows.Scan(
			&s.ID, &s.UserID, &s.PaymentMethodID,
			&s.Amount, &s.FreqVal, &s.FreqUnit,
			&s.CreatedAt, &s.UpdatedAt,
		)

		if err != nil {
			return nil, errors.New("error scanning subscription")
		}

		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error iterating over subscriptions")
	}

	return subs, nil
}

func (ds *dataSource) Delete(sub *entity.Subscription) error {
	_, err := ds.db.Pool.Exec("DELETE FROM subscriptions WHERE id = ?", sub.ID)
	if err != nil {
		return errors.New("error deleting subscription")
	}

	return nil
}
