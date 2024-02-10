package paymentmethod

import (
	"fmt"
	"net/http"
	"strconv"

	"harvest/bean/internal/entity/model"
	"harvest/bean/internal/entity/viewmodel"

	"harvest/bean/internal/usecase/paymentmethod"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	paymentMethods paymentmethod.UseCase

	view interfaces.CreatePaymentMethodView
}

func New(
	pm paymentmethod.UseCase,
	view interfaces.CreatePaymentMethodView,
) http.Handler {
	return &handler{
		paymentMethods: pm,
		view:           view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.render(w, &viewmodel.CreatePaymentMethodViewData{})
		return
	}

	if r.Method == http.MethodPost {
		h.post(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Method not allowed")
}

func (h *handler) post(w http.ResponseWriter, r *http.Request) {
	form := getFormData(r)

	_, err := h.paymentMethods.Create(
		"10000000-0000-0000-0000-000000000001",
		form.Label,
		form.Last4,
		model.PaymentMethodBrand(form.Brand),
		form.ExpMonth,
		form.ExpYear,
	)

	if err != nil {
		h.render(w, &viewmodel.CreatePaymentMethodViewData{
			Error: err.Error(),
			Form:  form,
		})
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *handler) render(w http.ResponseWriter, data *viewmodel.CreatePaymentMethodViewData) {
	err := h.view.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}

func getFormData(r *http.Request) viewmodel.CreatePaymentMethodForm {
	formData := viewmodel.CreatePaymentMethodForm{}

	if label := r.FormValue("label"); label != "" {
		formData.Label = label
	}

	if last4 := r.FormValue("last4"); last4 != "" {
		formData.Last4 = last4
	}

	if brand := r.FormValue("brand"); brand != "" {
		formData.Brand = brand
	}

	if expMonth := r.FormValue("exp_month"); expMonth != "" {
		expMonth, _ := strconv.Atoi(expMonth)

		formData.ExpMonth = expMonth
	}

	if expYear := r.FormValue("exp_year"); expYear != "" {
		expYear, _ := strconv.Atoi(expYear)

		formData.ExpYear = expYear
	}

	return formData
}
