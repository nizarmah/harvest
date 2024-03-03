package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	estimatorUC "github.com/whatis277/harvest/bean/internal/usecase/estimator"
	membershipUC "github.com/whatis277/harvest/bean/internal/usecase/membership"
	passwordlessUC "github.com/whatis277/harvest/bean/internal/usecase/passwordless"
	paymentMethodUC "github.com/whatis277/harvest/bean/internal/usecase/paymentmethod"
	subscriptionUC "github.com/whatis277/harvest/bean/internal/usecase/subscription"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/paymentmethod"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/subscription"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
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
	authVD "github.com/whatis277/harvest/bean/internal/driver/view/auth"
	landingVD "github.com/whatis277/harvest/bean/internal/driver/view/landing"
	paymentMethodVD "github.com/whatis277/harvest/bean/internal/driver/view/paymentmethod"
	subscriptionVD "github.com/whatis277/harvest/bean/internal/driver/view/subscription"
)

func main() {
	env, err := envAdapter.New()
	if err != nil {
		panic(
			fmt.Errorf("error reading env: %w", err),
		)
	}

	dbCtx, dbCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer dbCancel()

	db, err := postgres.New(dbCtx, &postgres.DSNBuilder{
		Host:     env.DB.Host,
		Port:     env.DB.Port,
		Name:     env.DB.Name,
		Username: env.DB.Username,
		Password: env.DB.Password,
		SSLMode:  env.DB.SSLMode,
	})
	if err != nil {
		panic(
			fmt.Errorf("error connecting db: %w", err),
		)
	}
	defer db.Close()

	fmt.Println("connected to db")

	cacheCtx, cacheCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cacheCancel()

	cache, err := redis.New(cacheCtx, &redis.Config{
		Host:        env.Cache.Host,
		Port:        env.Cache.Port,
		Username:    env.Cache.Username,
		Password:    env.Cache.Password,
		TLSDisabled: env.Cache.TLSDisabled,
	})
	if err != nil {
		panic(
			fmt.Errorf("error connecting cache: %w", err),
		)
	}
	defer cache.Close()

	fmt.Println("connected to cache")

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
	sessionRepo := sessionDS.New(cache, "session")
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
			fmt.Errorf("error creating landing view: %w", err),
		)
	}

	loginView, err := authVD.NewLogin(template.FS, template.LoginTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating login view: %w", err),
		)
	}

	signUpView, err := authVD.NewSignup(template.FS, template.SignUpTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating signup view: %w", err),
		)
	}

	homeView, err := appVD.NewHome(template.FS, template.HomeTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating home view: %w", err),
		)
	}

	renewPlanView, err := appVD.NewRenewPlan(template.FS, template.RenewPlanTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating renew plan view: %w", err),
		)
	}

	createPaymentMethodView, err := paymentMethodVD.NewCreate(template.FS, template.CreatePaymentMethodTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating create payment method view: %w", err),
		)
	}

	deletePaymentMethodView, err := paymentMethodVD.NewDelete(template.FS, template.DeletePaymentMethodTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating delete payment method view: %w", err),
		)
	}

	createSubscriptionView, err := subscriptionVD.NewCreate(template.FS, template.CreateSubscriptionTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating create subscription view: %w", err),
		)
	}

	deleteSubscriptionView, err := subscriptionVD.NewDelete(template.FS, template.DeleteSubscriptionTemplate)
	if err != nil {
		panic(
			fmt.Errorf("error creating delete subscription view: %w", err),
		)
	}

	baseController := base.Controller{}

	marketingController := marketing.Controller{
		LandingView: landingView,
	}

	authController := auth.Controller{
		BypassHTTPS: env.FeatureFlags.BypassHTTPS,

		Passwordless: passwordlessAuth,
		Memberships:  memberships,

		LoginView:  loginView,
		SignUpView: signUpView,
	}

	appController := app.Controller{
		Estimator:      estimator,
		PaymentMethods: paymentMethods,
		Memberships:    memberships,

		HomeView:      homeView,
		RenewPlanView: renewPlanView,
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
		AcceptTestEvents: env.BuyMeACoffee.AcceptTestEvents,
		WebhookSecret:    env.BuyMeACoffee.WebhookSecret,

		Passwordless: passwordlessAuth,
		Memberships:  memberships,
	}

	s := server.New(&server.Config{
		BaseHandler: baseController.ErrorHandler,
	})

	s.Route("GET /health", func(w http.ResponseWriter, r *http.Request) error {
		return nil
	})

	// Unauthenticated routes

	s.Route("GET /{$}", marketingController.LandingPage())

	s.Route("GET /auth/{id}/{password}", authController.Authorize())

	s.Route("GET /login", authController.LoginPage())
	s.Route("POST /login", authController.LoginForm())

	s.Route("GET /logout", authController.Logout())

	s.Route("GET /signup", authController.SignupPage())

	s.Route("POST /webhooks/buymeacoffee", bmcController.Webhook())

	// Authenticated routes

	s.Use(authController.Authenticate)

	s.Route("GET /renew-plan", appController.RenewPlanPage())

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

	fmt.Println("listening on :8080")
	s.Listen(":8080")
}
