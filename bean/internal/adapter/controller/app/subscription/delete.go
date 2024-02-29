package subscription

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/shared"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
)

func (c *Controller) DeletePage() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		subID := r.PathValue("id")
		if subID == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return nil
		}

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			auth.UnauthedUserRedirect(w, r)
			return nil
		}

		sub, err := c.Subscriptions.Get(ctx, session.UserID, subID)
		if err != nil {
			// FIXME: This should check for a specific error type
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return &base.HTTPError{
				Message: fmt.Sprintf(
					"subs: delete: error getting subscription: %v",
					err,
				),
			}
		}

		if sub == nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return nil
		}

		return c.renderDeleteView(w, &viewmodel.DeleteSubscriptionViewData{
			Subscription: shared.ToSubscriptionViewModel(sub),
		})
	}
}

func (c *Controller) DeleteForm() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		subID := r.FormValue("id")
		if subID == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return nil
		}

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			auth.UnauthedUserRedirect(w, r)
			return nil
		}

		err := c.Subscriptions.Delete(ctx, session.UserID, subID)
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return &base.HTTPError{
				Message: fmt.Sprintf(
					"subs: delete: error deleting subscription: %v",
					err,
				),
			}
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)

		return nil
	}
}

func (c *Controller) renderDeleteView(
	w http.ResponseWriter,
	data *viewmodel.DeleteSubscriptionViewData,
) error {
	err := c.DeleteView.Render(w, data)
	if err != nil {
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"subs: delete: error rendering delete view: %v",
				err,
			),
		}
	}

	return nil
}
