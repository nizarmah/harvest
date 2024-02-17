package paymentmethod

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	estimatorUC "github.com/whatis277/harvest/bean/internal/usecase/estimator"
	"github.com/whatis277/harvest/bean/internal/usecase/paymentmethod"

	"github.com/whatis277/harvest/bean/internal/adapter/handler/shared"
	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type deleteHandler struct {
	estimator      estimatorUC.UseCase
	paymentMethods paymentmethod.UseCase

	view interfaces.DeletePaymentMethodView
}

func newDeleteHandler(
	es estimatorUC.UseCase,
	pm paymentmethod.UseCase,
	view interfaces.DeletePaymentMethodView,
) http.Handler {
	return &deleteHandler{
		estimator:      es,
		paymentMethods: pm,
		view:           view,
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
	pmID := r.FormValue("id")
	if pmID == "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	pm, err := h.paymentMethods.Get(
		"10000000-0000-0000-0000-000000000001",
		pmID,
	)
	if err != nil || pm == nil {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	estimates := h.estimator.GetEstimates(pm.Subscriptions)

	h.render(w, &viewmodel.DeletePaymentMethodViewData{
		PaymentMethod: shared.ToPaymentMethodViewModel(pm, estimates),
	})
}

func (h *deleteHandler) post(w http.ResponseWriter, r *http.Request) {
	pmID := r.FormValue("id")
	if pmID == "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	h.paymentMethods.Delete(
		"10000000-0000-0000-0000-000000000001",
		pmID,
	)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *deleteHandler) render(w http.ResponseWriter, data *viewmodel.DeletePaymentMethodViewData) {
	err := h.view.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
