package landing

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase/paymentmethod"
	"harvest/bean/internal/usecase/userdash"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	paymentMethodUseCase paymentmethod.UseCase
	userDashUseCase      userdash.UseCase

	view interfaces.PaymentMethodsView
}

func New(
	pm paymentmethod.UseCase,
	ud userdash.UseCase,
	view interfaces.PaymentMethodsView,
) http.Handler {
	return &handler{
		paymentMethodUseCase: pm,
		userDashUseCase:      ud,
		view:                 view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methods, err := h.paymentMethodUseCase.List("10000000-0000-0000-0000-000000000001")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	err = h.view.Render(w, h.makeViewData(methods))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}

func (h *handler) makeViewData(paymentMethods []*entity.PaymentMethodWithSubscriptions) *entity.PaymentMethodsViewData {
	var methods = make([]entity.PaymentMethodViewData, 0, len(paymentMethods))
	for _, method := range paymentMethods {
		methods = append(methods, h.makeMethodViewData(method))
	}

	return &entity.PaymentMethodsViewData{
		PaymentMethods: methods,
	}
}

func (h *handler) makeMethodViewData(pm *entity.PaymentMethodWithSubscriptions) entity.PaymentMethodViewData {
	method, subs := pm.PaymentMethod, pm.Subscriptions

	estimates := h.userDashUseCase.GetEstimates(subs)

	monthlyDollars, monthlyCents := estimates.Monthly/100, estimates.Monthly%100
	yearlyDollars, yearlyCents := estimates.Yearly/100, estimates.Yearly%100

	return entity.PaymentMethodViewData{
		ID:       method.ID,
		Label:    method.Label,
		Last4:    method.Last4,
		Brand:    string(method.Brand),
		ExpMonth: method.ExpMonth,
		ExpYear:  method.ExpYear,

		MonthlyEstimate: fmt.Sprintf("$%d.%02d", monthlyDollars, monthlyCents),
		YearlyEstimate:  fmt.Sprintf("$%d.%02d", yearlyDollars, yearlyCents),
	}
}
