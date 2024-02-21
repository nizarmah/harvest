package subscription

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/shared"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
)

func (c *Controller) DeletePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subID := r.PathValue("id")
		if subID == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		session := auth.SessionFromContext(r.Context())

		sub, err := c.Subscriptions.Get(session.UserID, subID)
		if err != nil || sub == nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		c.renderDeleteView(w, &viewmodel.DeleteSubscriptionViewData{
			Subscription: shared.ToSubscriptionViewModel(sub),
		})
	}
}

func (c *Controller) DeleteForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subID := r.FormValue("id")
		if subID == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		session := auth.SessionFromContext(r.Context())

		c.Subscriptions.Delete(session.UserID, subID)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func (c *Controller) renderDeleteView(w http.ResponseWriter, data *viewmodel.DeleteSubscriptionViewData) {
	err := c.DeleteView.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
