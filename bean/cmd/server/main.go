package main

import (
	"fmt"

	estimatorUC "github.com/whatis277/harvest/bean/internal/usecase/estimator"
	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"
	paymentMethodUC "github.com/whatis277/harvest/bean/internal/usecase/paymentmethod"
	subscriptionUC "github.com/whatis277/harvest/bean/internal/usecase/subscription"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/marketing"
	envAdapter "github.com/whatis277/harvest/bean/internal/adapter/env"
	homeHandler "github.com/whatis277/harvest/bean/internal/adapter/handler/home"
	paymentMethodHandler "github.com/whatis277/harvest/bean/internal/adapter/handler/paymentmethod"
	subscriptionHandler "github.com/whatis277/harvest/bean/internal/adapter/handler/subscription"

	"github.com/whatis277/harvest/bean/internal/driver/bcrypt"
	paymentMethodDS "github.com/whatis277/harvest/bean/internal/driver/datasource/paymentmethod"
	sessionDS "github.com/whatis277/harvest/bean/internal/driver/datasource/session"
	subscriptionDS "github.com/whatis277/harvest/bean/internal/driver/datasource/subscription"
	tokenDS "github.com/whatis277/harvest/bean/internal/driver/datasource/token"
	userDS "github.com/whatis277/harvest/bean/internal/driver/datasource/user"
	"github.com/whatis277/harvest/bean/internal/driver/postgres"
	"github.com/whatis277/harvest/bean/internal/driver/redis"
	"github.com/whatis277/harvest/bean/internal/driver/server"
	"github.com/whatis277/harvest/bean/internal/driver/smtp"
	"github.com/whatis277/harvest/bean/internal/driver/template"
	homeVD "github.com/whatis277/harvest/bean/internal/driver/view/home"
	landingVD "github.com/whatis277/harvest/bean/internal/driver/view/landing"
	loginVD "github.com/whatis277/harvest/bean/internal/driver/view/login"
	paymentMethodVD "github.com/whatis277/harvest/bean/internal/driver/view/paymentmethod"
	subscriptionVD "github.com/whatis277/harvest/bean/internal/driver/view/subscription"
)

func main() {
	env, err := envAdapter.New()
	if err != nil {
		panic(
			fmt.Errorf("error reading env: %v", err),
		)
	}

	db, err := postgres.New(&postgres.DSNBuilder{
		Host:     env.DB.Host,
		Port:     env.DB.Port,
		Name:     env.DB.Name,
		Username: env.DB.Username,
		Password: env.DB.Password,
		SSLMode:  "disable",
	})
	if err != nil {
		panic(
			fmt.Errorf("error connecting db: %v", err),
		)
	}
	defer db.Close()

	cache, err := redis.New(&redis.Config{
		Host:     env.Cache.Host,
		Port:     env.Cache.Port,
		Username: env.Cache.Username,
		Password: env.Cache.Password,
	})
	if err != nil {
		panic(
			fmt.Errorf("error connecting cache: %v", err),
		)
	}
	defer cache.Close()

	hasher := bcrypt.New()
	emailer := smtp.New(&smtp.Config{
		Host:     env.SMTP.Host,
		Port:     env.SMTP.Port,
		Username: env.SMTP.Username,
		Password: env.SMTP.Password,
	})

	estimator := estimatorUC.UseCase{}

	paymentMethodRepo := paymentMethodDS.New(db)
	paymentMethods := paymentMethodUC.UseCase{
		PaymentMethods: paymentMethodRepo,
	}

	subscriptionRepo := subscriptionDS.New(db)
	subscriptions := subscriptionUC.UseCase{
		Subscriptions:  subscriptionRepo,
		PaymentMethods: paymentMethodRepo,
	}

	tokenRepo := tokenDS.New(db)
	userRepo := userDS.New(db)
	sessionRepo := sessionDS.New(cache)
	passwordlessAuth := passwordless.UseCase{
		Sender:   "Bean <support@whatisbean.com>",
		BaseURL:  env.BaseURL,
		Users:    userRepo,
		Tokens:   tokenRepo,
		Sessions: sessionRepo,
		Hasher:   hasher,
		Emailer:  emailer,
	}

	landingView, err := landingVD.New(template.FS, template.LandingTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating landing view: %v", err),
		)
	}

	loginView, err := loginVD.New(template.FS, template.LoginTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating login view: %v", err),
		)
	}

	homeView, err := homeVD.New(template.FS, template.HomeTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating home view: %v", err),
		)
	}

	createPaymentMethodView, err := paymentMethodVD.NewCreate(template.FS, template.CreatePaymentMethodTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating create payment method view: %v", err),
		)
	}

	deletePaymentMethodView, err := paymentMethodVD.NewDelete(template.FS, template.DeletePaymentMethodTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating delete payment method view: %v", err),
		)
	}

	createSubscriptionView, err := subscriptionVD.NewCreate(template.FS, template.CreateSubscriptionTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating create subscription view: %v", err),
		)
	}

	deleteSubscriptionView, err := subscriptionVD.NewDelete(template.FS, template.DeleteSubscriptionTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating delete subscription view: %v", err),
		)
	}

	marketingController := marketing.Controller{
		LandingView: landingView,
	}

	authControler := auth.Controller{
		Passwordless: passwordlessAuth,

		LoginView: loginView,
	}

	s := server.New()

	s.Route("/", marketingController.LandingPage())

	s.Route("GET /auth/{id}/{password}", authControler.Authorize())

	s.Route("GET /get-started", authControler.LoginPage())
	s.Route("POST /get-started", authControler.LoginForm())

	s.Route("/home", homeHandler.New(
		estimator,
		paymentMethods,
		homeView,
	))

	paymentMethodCRUD := paymentMethodHandler.New(
		estimator,
		paymentMethods,
		createPaymentMethodView,
		deletePaymentMethodView,
	)

	s.Route("/cards/new", paymentMethodCRUD.Create)
	s.Route("/cards/del", paymentMethodCRUD.Delete)

	subscriptionsCRUD := subscriptionHandler.New(
		subscriptions,
		createSubscriptionView,
		deleteSubscriptionView,
	)

	s.Route("/subs/new", subscriptionsCRUD.Create)
	s.Route("/subs/del", subscriptionsCRUD.Delete)

	s.Listen(":8080")
}
