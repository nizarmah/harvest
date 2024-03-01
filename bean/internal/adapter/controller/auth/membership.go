package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
)

func (c *Controller) CheckMembership(next base.HTTPHandler) base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		session := SessionFromContext(ctx)
		if session == nil {
			UnauthedUserRedirect(w, r)
			return nil
		}

		isMember, err := c.Memberships.CheckByID(ctx, session.UserID)
		if err != nil {
			return &base.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"auth: check-membership: error checking membership: %v",
					err,
				),
			}
		}

		if !isMember {
			http.Redirect(w, r, "/renew-plan", http.StatusSeeOther)
			return nil
		}

		return next(w, r)
	}
}
