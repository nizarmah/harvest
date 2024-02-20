package paymentmethod

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/whatis277/harvest/bean/internal/entity/model"
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"
)

func (c *Controller) CreatePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.renderCreateView(w, &viewmodel.CreatePaymentMethodViewData{})
	}
}

func (c *Controller) CreateForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := getCreateFormData(r)

		_, err := c.PaymentMethods.Create(
			"10000000-0000-0000-0000-000000000001",
			form.Label,
			form.Last4,
			model.PaymentMethodBrand(form.Brand),
			form.ExpMonth,
			form.ExpYear,
		)

		if err != nil {
			c.renderCreateView(w, &viewmodel.CreatePaymentMethodViewData{
				Error: err.Error(),
				Form:  form,
			})
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func (c *Controller) renderCreateView(w http.ResponseWriter, data *viewmodel.CreatePaymentMethodViewData) {
	err := c.CreateView.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}

func getCreateFormData(r *http.Request) viewmodel.CreatePaymentMethodForm {
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
