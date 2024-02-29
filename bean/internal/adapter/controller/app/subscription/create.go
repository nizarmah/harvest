package subscription

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/whatis277/harvest/bean/internal/entity/model"
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
)

func (c *Controller) CreatePage() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		pmID := r.PathValue("pm_id")
		if pmID == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return &base.HTTPError{
				Message: "subs: create: no payment method id provided",
			}
		}

		return c.renderCreateView(w, &viewmodel.CreateSubscriptionViewData{
			Form: viewmodel.CreateSubscriptionForm{
				PaymentMethodID: pmID,
			},
		})
	}
}

func (c *Controller) CreateForm() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		form := getCreateFormData(r)

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			auth.UnauthedUserRedirect(w, r)
			return nil
		}

		_, err := c.Subscriptions.Create(
			ctx,
			session.UserID,
			form.PaymentMethodID,
			form.Label,
			form.Provider,
			inCents(form.Amount),
			form.Interval,
			model.SubscriptionPeriod(form.Period),
		)

		if err != nil {
			// FIXME: This should check for a specific error type
			return c.renderCreateView(w, &viewmodel.CreateSubscriptionViewData{
				Error: err.Error(),
				Form:  form,
			})
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)

		return nil
	}
}

func (c *Controller) renderCreateView(
	w http.ResponseWriter,
	data *viewmodel.CreateSubscriptionViewData,
) error {
	err := c.CreateView.Render(w, data)
	if err != nil {
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"subs: create: error rendering create view: %v",
				err,
			),
		}
	}

	return nil
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
		amount, err := strconv.ParseFloat(amount, 32)
		if err == nil {
			formData.Amount = float32(amount)
		}
	}

	if interval := r.FormValue("interval"); interval != "" {
		interval, err := strconv.Atoi(interval)
		if err == nil {
			formData.Interval = interval
		}
	}

	if period := r.FormValue("period"); period != "" {
		formData.Period = period
	}

	return formData
}

func inCents(amount float32) int {
	return int(amount * 100)
}
