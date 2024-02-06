package landing

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase/paymentmethod"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	usecase paymentmethod.UseCase

	view interfaces.PaymentMethodsView
}

func New(
	usecase paymentmethod.UseCase,
	view interfaces.PaymentMethodsView,
) http.Handler {
	return &handler{
		usecase: usecase,
		view:    view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methods, err := h.usecase.List("10000000-0000-0000-0000-000000000001")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	err = h.view.Render(w, &entity.PaymentMethodsViewData{
		PaymentMethods: methods,
	})
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
