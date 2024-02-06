package landing

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase/subscription"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	usecase subscription.UseCase

	view interfaces.SubscriptionsView
}

func New(
	usecase subscription.UseCase,
	view interfaces.SubscriptionsView,
) http.Handler {
	return &handler{
		usecase: usecase,
		view:    view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	subs, err := h.usecase.List("10000000-0000-0000-0000-000000000001")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	err = h.view.Render(w, &entity.SubscriptionsViewData{
		Subscriptions: subs,
	})
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
