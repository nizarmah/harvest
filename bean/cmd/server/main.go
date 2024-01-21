package main

import (
	"fmt"
	"time"

	"harvest/bean/internal/entity"

	envAdapter "harvest/bean/internal/adapter/env"

	"harvest/bean/internal/driver/database"
	subscriptionDS "harvest/bean/internal/driver/datasource/subscription"
)

func main() {
	env, err := envAdapter.New()
	if err != nil {
		panic(
			fmt.Errorf("error reading env: %v", err),
		)
	}

	db, err := database.New(&database.DSNBuilder{
		Host:        env.DB.Host,
		Name:        env.DB.Name,
		Username:    env.DB.Username,
		Password:    env.DB.Password,
		Tls:         true,
		Interpolate: true,
		ParseTime:   true,
	})
	if err != nil {
		panic(
			fmt.Errorf("error connecting db: %v", err),
		)
	}
	defer db.Close()

	u := &entity.User{
		ID:        "test-user",
		Email:     "test-user@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testSubscriptions(db, u)
}

func testSubscriptions(db *database.DB, u *entity.User) {
	subs := subscriptionDS.New(db)

	_, err := subs.Create(&entity.Subscription{
		UserID:          u.ID,
		PaymentMethodID: 1,
		Amount:          100,
		FreqVal:         1,
		FreqUnit:        "month",
	})
	if err != nil {
		fmt.Println("error creating subscription 1: ", err)
		return
	}

	_, err = subs.Create(&entity.Subscription{
		UserID:          u.ID,
		PaymentMethodID: 1,
		Amount:          200,
		FreqVal:         6,
		FreqUnit:        "month",
	})
	if err != nil {
		fmt.Println("error creating subscription 2: ", err)
		return
	}

	slice, err := subs.FindByUserId(u.ID)
	if err != nil {
		fmt.Println("error finding subscriptions: ", err)
		return
	}

	for _, s := range slice {
		fmt.Println(
			"subscription found: ",
			s.ID, s.UserID, s.PaymentMethodID,
			s.Amount, s.FreqVal, s.FreqUnit,
			s.CreatedAt, s.UpdatedAt,
		)

		err = subs.Delete(s)
		if err != nil {
			fmt.Println("error deleting subscription: ", err)
			return
		}
	}
}
