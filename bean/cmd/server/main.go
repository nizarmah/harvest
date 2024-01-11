package main

import (
	"fmt"

	"harvest/bean/internal/entity"

	envAdapter "harvest/bean/internal/adapter/env"

	"harvest/bean/internal/driver/crypto"
	"harvest/bean/internal/driver/database"
	paymentMethodDS "harvest/bean/internal/driver/datasource/paymentmethod"
	subscriptionDS "harvest/bean/internal/driver/datasource/subscription"
	tokenDS "harvest/bean/internal/driver/datasource/token"
	userDS "harvest/bean/internal/driver/datasource/user"
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

	u := createUserIfNotExists(db)
	if u == nil {
		fmt.Println("user not found")
		return
	}

	testLoginTokens(db, u)
	testPaymentMethods(db, u)
	testSubscriptions(db, u)
}

func createUserIfNotExists(db *database.DB) *entity.User {
	users := userDS.New(db)

	u, _ := users.FindByEmail("test@gmail.com")
	if u != nil {
		fmt.Println("user found by email: ", u.ID, u.Email, u.CreatedAt, u.UpdatedAt)
		return u
	}

	u, err := users.Create(&entity.User{Email: "test@gmail.com"})
	if u != nil {
		fmt.Println("user created: ", u.ID, u.Email, u.CreatedAt, u.UpdatedAt)
	} else {
		fmt.Println("user not created: ", err)
		return nil
	}

	u, err = users.FindById(u.ID)
	if u != nil {
		fmt.Println("user found by id: ", u.ID, u.Email, u.CreatedAt, u.UpdatedAt)
	} else {
		fmt.Println("user not found by id: ", err)
	}

	return u
}

func testLoginTokens(db *database.DB, u *entity.User) {
	tokens := tokenDS.New(db)
	hasher := crypto.New()

	hash1, err := hasher.Hash("123")
	if err != nil {
		fmt.Println("error hashing token 1: ", err)
		return
	}

	hash2, err := hasher.Hash("456")
	if err != nil {
		fmt.Println("error hashing token 2: ", err)
		return
	}

	err = tokens.Create(&entity.LoginToken{Email: u.Email, HashedToken: hash1})
	if err != nil {
		fmt.Println("error creating token: ", err)
		return
	}

	t, err := tokens.FindUnexpired(&entity.LoginToken{Email: u.Email, HashedToken: hash1})
	if t != nil {
		fmt.Println("token found: ", t.ID, t.Email, string(t.HashedToken), t.CreatedAt, t.ExpiresAt)
	} else {
		fmt.Println("token not found: ", err)
	}

	err = tokens.Create(&entity.LoginToken{Email: u.Email, HashedToken: hash2})
	if err != nil {
		fmt.Println("error overwriting token: ", err)
	}

	t, err = tokens.FindUnexpired(&entity.LoginToken{Email: u.Email, HashedToken: hash2})
	if t != nil {
		fmt.Println("token found: ", t.ID, t.Email, string(t.HashedToken), t.CreatedAt, t.ExpiresAt)
	} else {
		fmt.Println("token not found: ", err)
		return
	}

	err = tokens.Delete(t)
	if err != nil {
		fmt.Println("error deleting token: ", err)
	}
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
