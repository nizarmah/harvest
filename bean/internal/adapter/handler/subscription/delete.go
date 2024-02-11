package subscription

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity/viewmodel"

	"harvest/bean/internal/usecase/subscription"

	"harvest/bean/internal/adapter/handler/shared"
	"harvest/bean/internal/adapter/interfaces"
)

type deleteHandler struct {
	subscriptions subscription.UseCase

	view interfaces.DeleteSubscriptionView
}

func newDeleteHandler(
	sub subscription.UseCase,
	view interfaces.DeleteSubscriptionView,
) http.Handler {
	return &deleteHandler{
		subscriptions: sub,
		view:          view,
	}
}

func (h *deleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.get(w, r)
		return
	}

	if r.Method == http.MethodPost {
		h.post(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Method not allowed")
}

func (h *deleteHandler) get(w http.ResponseWriter, r *http.Request) {
	subID := r.FormValue("id")
	if subID == "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	sub, err := h.subscriptions.Get(
		"10000000-0000-0000-0000-000000000001",
		subID,
	)
	if err != nil || sub == nil {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	h.render(w, &viewmodel.DeleteSubscriptionViewData{
		Subscription: shared.ToSubscriptionViewModel(sub),
	})
}

func (h *deleteHandler) post(w http.ResponseWriter, r *http.Request) {
	subID := r.FormValue("id")
	if subID == "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	h.subscriptions.Delete(
		"10000000-0000-0000-0000-000000000001",
		subID,
	)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *deleteHandler) render(w http.ResponseWriter, data *viewmodel.DeleteSubscriptionViewData) {
	err := h.view.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
