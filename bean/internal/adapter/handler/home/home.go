package home

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity/viewmodel"

	estimatorUC "harvest/bean/internal/usecase/estimator"
	"harvest/bean/internal/usecase/paymentmethod"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	estimator      estimatorUC.UseCase
	paymentMethods paymentmethod.UseCase

	view interfaces.PaymentMethodsView
}

func New(
	es estimatorUC.UseCase,
	pm paymentmethod.UseCase,
	view interfaces.PaymentMethodsView,
) http.Handler {
	return &handler{
		estimator:      es,
		paymentMethods: pm,
		view:           view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methods, err := h.paymentMethods.List("10000000-0000-0000-0000-000000000001")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	methodsVM := make([]viewmodel.PaymentMethod, 0, len(methods))
	totalMonthly, totalYearly := 0, 0
	for _, method := range methods {
		estimates := h.estimator.GetEstimates(method.Subscriptions)

		totalMonthly += estimates.Monthly
		totalYearly += estimates.Yearly

		methodsVM = append(methodsVM, toPaymentMethodViewModel(method, estimates))
	}

	err = h.view.Render(w, &viewmodel.HomeViewData{
		PaymentMethods:  methodsVM,
		MonthlyEstimate: toDollarsString(totalMonthly),
		YearlyEstimate:  toDollarsString(totalYearly),
	})
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
