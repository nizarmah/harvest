package subscription

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity/viewmodel"

	"harvest/bean/internal/usecase/subscription"

	"harvest/bean/internal/adapter/interfaces"
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

	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Method not allowed")
}

func (h *createHandler) render(w http.ResponseWriter, data *viewmodel.CreateSubscriptionViewData) {
	err := h.view.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
