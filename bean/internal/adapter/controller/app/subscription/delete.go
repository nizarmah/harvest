package subscription

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/shared"
)

func (c *Controller) DeletePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subID := r.PathValue("id")
		if subID == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		sub, err := c.Subscriptions.Get(
			"10000000-0000-0000-0000-000000000001",
			subID,
		)
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

		c.Subscriptions.Delete(
			"10000000-0000-0000-0000-000000000001",
			subID,
		)

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
