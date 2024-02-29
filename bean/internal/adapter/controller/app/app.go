package app

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
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

	HomeView      interfaces.HomeView
	RenewPlanView interfaces.RenewPlanView
}

func (c *Controller) HomePage() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			return auth.NewUnauthorizedError(
				"app: home: user has no session",
			)
		}

		methods, err := c.PaymentMethods.List(ctx, session.UserID)
		if err != nil {
			// FIXME: This should check for a specific error type
			return &model.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"app: home: error listing payment methods: %v",
					err,
				),
			}
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
			return &model.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"app: home: error rendering home view: %v",
					err,
				),
			}
		}

		return nil
	}
}

func (c *Controller) RenewPlanPage() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			return auth.NewUnauthorizedError(
				"app: renew-plan: user has no session",
			)
		}

		isMember, _ := c.Memberships.CheckByID(ctx, session.UserID)
		if isMember {
			return &model.HTTPError{
				Status:       http.StatusFound,
				RedirectPath: "/home",

				Message: "app: renew-plan: user is already a member",
			}
		}

		err := c.RenewPlanView.Render(w, nil)
		if err != nil {
			return &model.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"app: renew-plan: error rendering renew plan view: %v",
					err,
				),
			}
		}

		return nil
	}
}
