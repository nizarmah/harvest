package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

func (c *Controller) Logout() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		c.cleanupSessionToken(w)

		session := SessionFromContext(ctx)
		if session == nil {
			return NewUnauthorizedError(
				"auth: logout: user has no session",
			)
		}

		err := c.Passwordless.Logout(ctx, session)
		if err != nil {
			return NewUnauthorizedError(
				fmt.Sprintf(
					"auth: logout: error logging out user: %v",
					err,
				),
			)
		}

		http.Redirect(
			w,
			r,
			defaultUnauthedRedirectPath,
			defaultObscureStatus,
		)

		return nil
	}
}
