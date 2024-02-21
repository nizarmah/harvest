package subscription

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/whatis277/harvest/bean/internal/entity/model"
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"
)

func (c *Controller) CreatePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pmID := r.PathValue("pm_id")

		c.renderCreateView(w, &viewmodel.CreateSubscriptionViewData{
			Form: viewmodel.CreateSubscriptionForm{
				PaymentMethodID: pmID,
			},
		})
	}
}

func (c *Controller) CreateForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := getCreateFormData(r)

		_, err := c.Subscriptions.Create(
			"10000000-0000-0000-0000-000000000001",
			form.PaymentMethodID,
			form.Label,
			form.Provider,
			inCents(form.Amount),
			form.Interval,
			model.SubscriptionPeriod(form.Period),
		)

		if err != nil {
			c.renderCreateView(w, &viewmodel.CreateSubscriptionViewData{
				Error: err.Error(),
				Form:  form,
			})
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func (c *Controller) renderCreateView(w http.ResponseWriter, data *viewmodel.CreateSubscriptionViewData) {
	err := c.CreateView.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}

func getCreateFormData(r *http.Request) viewmodel.CreateSubscriptionForm {
	formData := viewmodel.CreateSubscriptionForm{}

	if pmID := r.FormValue("pm_id"); pmID != "" {
		formData.PaymentMethodID = pmID
	}

	if label := r.FormValue("label"); label != "" {
		formData.Label = label
	}

	if provider := r.FormValue("provider"); provider != "" {
		formData.Provider = provider
	}

	if amount := r.FormValue("amount"); amount != "" {
		amount, _ := strconv.ParseFloat(amount, 32)

		formData.Amount = float32(amount)
	}

	if interval := r.FormValue("interval"); interval != "" {
		interval, _ := strconv.Atoi(interval)

		formData.Interval = interval
	}

	if period := r.FormValue("period"); period != "" {
		formData.Period = period
	}

	return formData
}

func inCents(amount float32) int {
	return int(amount * 100)
}
