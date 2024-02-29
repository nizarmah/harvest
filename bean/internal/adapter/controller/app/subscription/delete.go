package subscription

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/app/shared"
	"github.com/whatis277/harvest/bean/internal/adapter/controller/auth"
)

func (c *Controller) DeletePage() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		subID := r.PathValue("id")
		if subID == "" {
			return &model.HTTPError{
				Status:       http.StatusSeeOther,
				RedirectPath: "/home",

				Message: "subs: delete: no id provided",
			}
		}

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			return auth.NewUnauthorizedError(
				"subs: delete: user has no session",
			)
		}

		sub, err := c.Subscriptions.Get(ctx, session.UserID, subID)
		if err != nil || sub == nil {
			// FIXME: This should check for a specific error type
			return &model.HTTPError{
				Status:       http.StatusSeeOther,
				RedirectPath: "/home",

				Message: fmt.Sprintf(
					"subs: delete: error getting subscription: %v",
					err,
				),
			}
		}

		return c.renderDeleteView(w, &viewmodel.DeleteSubscriptionViewData{
			Subscription: shared.ToSubscriptionViewModel(sub),
		})
	}
}

func (c *Controller) DeleteForm() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		subID := r.FormValue("id")
		if subID == "" {
			return &model.HTTPError{
				Status:       http.StatusSeeOther,
				RedirectPath: "/home",

				Message: "subs: delete: no id provided",
			}
		}

		ctx := r.Context()

		session := auth.SessionFromContext(ctx)
		if session == nil {
			return auth.NewUnauthorizedError(
				"subs: delete: user has no session",
			)
		}

		err := c.Subscriptions.Delete(ctx, session.UserID, subID)
		if err != nil {
			return &model.HTTPError{
				Status:       http.StatusSeeOther,
				RedirectPath: "/home",

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
		return &model.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"subs: delete: error rendering delete view: %v",
				err,
			),
		}
	}

	return nil
}
