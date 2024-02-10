package paymentmethod

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	view interfaces.CreatePaymentMethodView
}

func New(view interfaces.CreatePaymentMethodView) http.Handler {
	return &handler{
		view: view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.view.Render(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
