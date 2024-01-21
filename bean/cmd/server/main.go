package main

import (
	"fmt"
	"time"

	"harvest/bean/internal/entity"

	envAdapter "harvest/bean/internal/adapter/env"

	"harvest/bean/internal/driver/database"
	paymentMethodDS "harvest/bean/internal/driver/datasource/paymentmethod"
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

	testPaymentMethods(db, u)
	testSubscriptions(db, u)
}

func testPaymentMethods(db *database.DB, u *entity.User) {
	methods := paymentMethodDS.New(db)

	_, err := methods.Create(&entity.PaymentMethod{
		UserID:    u.ID,
		Label:     "test",
		Last4:     "1234",
		Brand:     "visa",
		ExpMonth:  7,
		ExpYear:   2028,
		IsDefault: true,
	})
	if err != nil {
		fmt.Println("error creating payment method 1: ", err)
		return
	}

	_, err = methods.Create(&entity.PaymentMethod{
		UserID:    u.ID,
		Label:     "test",
		Last4:     "5678",
		Brand:     "visa",
		ExpMonth:  7,
		ExpYear:   2028,
		IsDefault: false,
	})
	if err != nil {
		fmt.Println("error creating payment method 2: ", err)
		return
	}

	slice, err := methods.FindByUserId(u.ID)
	if err != nil {
		fmt.Println("error finding payment methods: ", err)
		return
	}

	for _, m := range slice {
		fmt.Println(
			"payment method found: ",
			m.ID, m.UserID,
			m.Label, m.Last4, m.Brand, m.ExpMonth, m.ExpYear,
			m.IsDefault,
			m.CreatedAt, m.UpdatedAt,
		)

		err = methods.Delete(m)
		if err != nil {
			fmt.Println("error deleting payment method: ", err)
			return
		}
	}
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
