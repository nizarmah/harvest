package paymentmethod

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/shared"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
)

func (c *Controller) DeletePage() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		if id == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return nil
		}

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			auth.UnauthedUserRedirect(w, r)
			return nil
		}

		pm, err := c.PaymentMethods.Get(ctx, session.UserID, id)
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return &base.HTTPError{
				Message: fmt.Sprintf(
					"pms: delete: error getting payment method: %v",
					err,
				),
			}
		}

		if pm == nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return nil
		}

		estimates := c.Estimator.GetEstimates(pm.Subscriptions)

		return c.renderDeleteView(w, &viewmodel.DeletePaymentMethodViewData{
			PaymentMethod: shared.ToPaymentMethodViewModel(pm, estimates),
		})
	}
}

func (c *Controller) DeleteForm() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.FormValue("id")
		if id == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return nil
		}

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			auth.UnauthedUserRedirect(w, r)
			return nil
		}

		err := c.PaymentMethods.Delete(ctx, session.UserID, id)
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return &base.HTTPError{
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
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"pms: delete: error rendering delete view: %v",
				err,
			),
		}
	}

	return nil
}
