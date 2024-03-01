package paymentmethod

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
		return c.renderCreateView(w, &viewmodel.CreatePaymentMethodViewData{})
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

		_, err := c.PaymentMethods.Create(
			ctx,
			session.UserID,
			form.Label,
			form.Last4,
			model.PaymentMethodBrand(form.Brand),
			form.ExpMonth,
			form.ExpYear,
		)

		switch err.(type) {
		case nil:
			break

		case model.UserInputError:
			return c.renderCreateView(w, &viewmodel.CreatePaymentMethodViewData{
				Error: err.Error(),
				Form:  form,
			})

		default:
			renderErr := c.renderCreateView(w, &viewmodel.CreatePaymentMethodViewData{
				Error: "Something went wrong",
				Form:  form,
			})
			return &base.HTTPError{
				Message: fmt.Sprintf(
					("pms: create: error creating payment method: %v |" +
						" pms: create: error rendering view: %v"),
					err, renderErr,
				),
			}
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)

		return nil
	}
}

func (c *Controller) renderCreateView(
	w http.ResponseWriter,
	data *viewmodel.CreatePaymentMethodViewData,
) error {
	err := c.CreateView.Render(w, data)
	if err != nil {
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"pms: create: error rendering create view: %v",
				err,
			),
		}
	}

	return nil
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
