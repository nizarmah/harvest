package main

import (
	"fmt"

	estimatorUC "github.com/whatis277/harvest/bean/internal/usecase/estimator"
	membershipUC "github.com/whatis277/harvest/bean/internal/usecase/membership"
	passwordlessUC "github.com/whatis277/harvest/bean/internal/usecase/passwordless"
	paymentMethodUC "github.com/whatis277/harvest/bean/internal/usecase/paymentmethod"
	subscriptionUC "github.com/whatis277/harvest/bean/internal/usecase/subscription"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/paymentmethod"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/subscription"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/marketing"
	envAdapter "github.com/whatis277/harvest/bean/internal/adapter/env"

	"github.com/whatis277/harvest/bean/internal/driver/bcrypt"
	"github.com/whatis277/harvest/bean/internal/driver/buymeacoffee"
	membershipDS "github.com/whatis277/harvest/bean/internal/driver/datasource/membership"
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
	appVD "github.com/whatis277/harvest/bean/internal/driver/view/app"
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
	passwordlessAuth := passwordlessUC.UseCase{
		Sender:   "Bean <support@whatisbean.com>",
		BaseURL:  env.BaseURL,
		Users:    userRepo,
		Tokens:   tokenRepo,
		Sessions: sessionRepo,
		Hasher:   hasher,
		Emailer:  emailer,
	}

	membershipRepo := membershipDS.New(db)
	memberships := membershipUC.UseCase{
		Bypass:      env.FeatureFlags.BypassMembership,
		Users:       userRepo,
		Memberships: membershipRepo,
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

	homeView, err := appVD.NewHome(template.FS, template.HomeTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating home view: %v", err),
		)
	}

	onboardingView, err := appVD.NewOnboarding(template.FS, template.OnboardingTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating onboarding view: %v", err),
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

	authController := auth.Controller{
		BypassHTTPS: env.FeatureFlags.BypassHTTPS,

		Passwordless: passwordlessAuth,
		Memberships:  memberships,

		LoginView: loginView,
	}

	appController := app.Controller{
		Estimator:      estimator,
		PaymentMethods: paymentMethods,
		Memberships:    memberships,

		HomeView:       homeView,
		OnboardingView: onboardingView,
	}

	pmsController := paymentmethod.Controller{
		Estimator:      estimator,
		PaymentMethods: paymentMethods,

		CreateView: createPaymentMethodView,
		DeleteView: deletePaymentMethodView,
	}

	subsController := subscription.Controller{
		Subscriptions: subscriptions,

		CreateView: createSubscriptionView,
		DeleteView: deleteSubscriptionView,
	}

	bmcController := buymeacoffee.Controller{
		WebhookSecret: env.BuyMeACoffee.WebhookSecret,

		Memberships: memberships,
	}

	s := server.New()

	// Unauthenticated routes

	s.Route("GET /{$}", marketingController.LandingPage())

	s.Route("GET /auth/{id}/{password}", authController.Authorize())

	s.Route("GET /get-started", authController.LoginPage())
	s.Route("POST /get-started", authController.LoginForm())

	s.Route("POST /webhooks/buymeacoffee", bmcController.Webhook())

	// Authenticated routes

	s.Use(authController.Authenticate)

	s.Route("GET /onboarding", appController.OnboardingPage())

	s.Route("GET /logout", authController.Logout())

	s.Use(authController.CheckMembership)

	s.Route("GET /home", appController.HomePage())

	s.Route("GET /cards/new", pmsController.CreatePage())
	s.Route("POST /cards/new", pmsController.CreateForm())

	s.Route("GET /cards/{id}/del", pmsController.DeletePage())
	s.Route("POST /cards/{id}/del", pmsController.DeleteForm())

	s.Route("GET /cards/{pm_id}/subs/new", subsController.CreatePage())
	s.Route("POST /cards/{pm_id}/subs/new", subsController.CreateForm())

	s.Route("GET /cards/{pm_id}/subs/{id}/del", subsController.DeletePage())
	s.Route("POST /cards/{pm_id}/subs/{id}/del", subsController.DeleteForm())

	s.Listen(":8080")
}
