package paymentmethod

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/shared"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
)

func (c *Controller) DeletePage() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		if id == "" {
			return &model.HTTPError{
				Status:       http.StatusSeeOther,
				RedirectPath: "/home",

				Message: "pms: delete: no id provided",
			}
		}

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			return auth.NewUnauthorizedError(
				"pms: delete: user has no session",
			)
		}

		pm, err := c.PaymentMethods.Get(ctx, session.UserID, id)
		if err != nil || pm == nil {
			return &model.HTTPError{
				Status:       http.StatusSeeOther,
				RedirectPath: "/home",

				Message: fmt.Sprintf(
					"pms: delete: error getting payment method: %v",
					err,
				),
			}
		}

		estimates := c.Estimator.GetEstimates(pm.Subscriptions)

		return c.renderDeleteView(w, &viewmodel.DeletePaymentMethodViewData{
			PaymentMethod: shared.ToPaymentMethodViewModel(pm, estimates),
		})
	}
}

func (c *Controller) DeleteForm() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.FormValue("id")
		if id == "" {
			return &model.HTTPError{
				Status:       http.StatusSeeOther,
				RedirectPath: "/home",

				Message: "pms: delete: no id provided",
			}
		}

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			return auth.NewUnauthorizedError(
				"pms: delete: user has no session",
			)
		}

		err := c.PaymentMethods.Delete(ctx, session.UserID, id)
		if err != nil {
			return &model.HTTPError{
				Status:       http.StatusSeeOther,
				RedirectPath: "/home",

				Message: fmt.Sprintf(
					"pms: delete: error deleting payment method: %v",
					err,
				),
			}
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)

		return nil
	}
}

func (c *Controller) renderDeleteView(
	w http.ResponseWriter,
	data *viewmodel.DeletePaymentMethodViewData,
) error {
	err := c.DeleteView.Render(w, data)
	if err != nil {
		return &model.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"pms: delete: error rendering delete view: %v",
				err,
			),
		}
	}

	return nil
}
