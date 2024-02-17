package subscription

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/whatis277/harvest/bean/internal/entity/model"
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/usecase/subscription"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type createHandler struct {
	subscriptions subscription.UseCase

	view interfaces.CreateSubscriptionView
}

func newCreateHandler(
	sub subscription.UseCase,
	view interfaces.CreateSubscriptionView,
) http.Handler {
	return &createHandler{
		subscriptions: sub,
		view:          view,
	}
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.render(w, &viewmodel.CreateSubscriptionViewData{})
		return
	}

	if r.Method == http.MethodPost {
		h.post(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Method not allowed")
}

func (h *createHandler) post(w http.ResponseWriter, r *http.Request) {
	form := getFormData(r)

	_, err := h.subscriptions.Create(
		"10000000-0000-0000-0000-000000000001",
		form.PaymentMethodID,
		form.Label,
		form.Provider,
		getAmountInCents(form.Amount),
		form.Interval,
		model.SubscriptionPeriod(form.Period),
	)

	if err != nil {
		h.render(w, &viewmodel.CreateSubscriptionViewData{
			Error: err.Error(),
			Form:  form,
		})
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *createHandler) render(w http.ResponseWriter, data *viewmodel.CreateSubscriptionViewData) {
	err := h.view.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}

func getFormData(r *http.Request) viewmodel.CreateSubscriptionForm {
	formData := viewmodel.CreateSubscriptionForm{}

	if pmID := r.FormValue("pm"); pmID != "" {
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

func getAmountInCents(amount float32) int {
	return int(amount * 100)
}
