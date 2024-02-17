package home

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	estimatorUC "github.com/whatis277/harvest/bean/internal/usecase/estimator"
	"github.com/whatis277/harvest/bean/internal/usecase/paymentmethod"

	"github.com/whatis277/harvest/bean/internal/adapter/handler/shared"
	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	estimator      estimatorUC.UseCase
	paymentMethods paymentmethod.UseCase

	view interfaces.HomeView
}

func New(
	es estimatorUC.UseCase,
	pm paymentmethod.UseCase,
	view interfaces.HomeView,
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

		methodsVM = append(methodsVM, shared.ToPaymentMethodViewModel(method, estimates))
	}

	err = h.view.Render(w, &viewmodel.HomeViewData{
		PaymentMethods:  methodsVM,
		MonthlyEstimate: shared.ToDollarsString(totalMonthly),
		YearlyEstimate:  shared.ToDollarsString(totalYearly),
	})
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
