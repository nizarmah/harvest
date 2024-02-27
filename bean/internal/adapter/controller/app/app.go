package app

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/usecase/estimator"
	"github.com/whatis277/harvest/bean/internal/usecase/membership"
	"github.com/whatis277/harvest/bean/internal/usecase/paymentmethod"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/shared"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	Estimator      estimator.UseCase
	PaymentMethods paymentmethod.UseCase
	Memberships    membership.UseCase

	HomeView       interfaces.HomeView
	OnboardingView interfaces.OnboardingView
}

func (c *Controller) HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := auth.SessionFromContext(r.Context())

		methods, err := c.PaymentMethods.List(session.UserID)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		methodsVM := make([]viewmodel.PaymentMethod, 0, len(methods))
		totalMonthly, totalYearly := 0, 0
		for _, method := range methods {
			estimates := c.Estimator.GetEstimates(method.Subscriptions)

			totalMonthly += estimates.Monthly
			totalYearly += estimates.Yearly

			methodsVM = append(methodsVM, shared.ToPaymentMethodViewModel(method, estimates))
		}

		err = c.HomeView.Render(w, &viewmodel.HomeViewData{
			PaymentMethods:  methodsVM,
			MonthlyEstimate: shared.ToDollarsString(totalMonthly),
			YearlyEstimate:  shared.ToDollarsString(totalYearly),
		})
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
	}
}

func (c *Controller) OnboardingPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := auth.SessionFromContext(r.Context())

		isMember, _ := c.Memberships.CheckByID(session.UserID)
		if isMember {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}

		err := c.OnboardingView.Render(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
	}
}
