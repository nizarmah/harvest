package paymentmethod

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/shared"
)

func (c *Controller) DeletePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		pm, err := c.PaymentMethods.Get(
			"10000000-0000-0000-0000-000000000001",
			id,
		)
		if err != nil || pm == nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		estimates := c.Estimator.GetEstimates(pm.Subscriptions)

		c.renderDeleteView(w, &viewmodel.DeletePaymentMethodViewData{
			PaymentMethod: shared.ToPaymentMethodViewModel(pm, estimates),
		})
	}
}

func (c *Controller) DeleteForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		if id == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		c.PaymentMethods.Delete(
			"10000000-0000-0000-0000-000000000001",
			id,
		)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func (c *Controller) renderDeleteView(w http.ResponseWriter, data *viewmodel.DeletePaymentMethodViewData) {
	err := c.DeleteView.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
