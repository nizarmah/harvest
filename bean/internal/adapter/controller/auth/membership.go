package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

func (c *Controller) CheckMembership(next model.HTTPHandler) model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		session := SessionFromContext(ctx)
		if session == nil {
			return NewUnauthorizedError(
				"auth: check-membership: user has no session",
			)
		}

		isMember, err := c.Memberships.CheckByID(ctx, session.UserID)
		if err != nil {
			// FIXME: This should check for a specific error type
			return &model.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"auth: check-membership: error checking membership: %v",
					err,
				),
			}
		}

		if !isMember {
			return &model.HTTPError{
				Status:       http.StatusFound,
				RedirectPath: "/renew-plan",

				Message: "auth: check-membership: user is not a member",
			}
		}

		return next(w, r)
	}
}
