package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
)

func (c *Controller) Logout() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		sessionToken := c.getSessionToken(r)
		if sessionToken == nil {
			http.Redirect(
				w,
				r,
				"/login",
				defaultObscureStatus,
			)
			return nil
		}

		ctx := r.Context()

		c.cleanupSessionToken(w)

		session := SessionFromContext(ctx)
		if session == nil {
			http.Redirect(
				w,
				r,
				"/login",
				defaultObscureStatus,
			)
			return nil
		}

		err := c.Passwordless.Logout(ctx, session)
		if err != nil {
			http.Redirect(
				w,
				r,
				"/login",
				defaultObscureStatus,
			)
			return &base.HTTPError{
				Message: fmt.Sprintf(
					"auth: logout: error logging out: %v",
					err,
				),
			}
		}

		http.Redirect(
			w,
			r,
			"/login",
			defaultObscureStatus,
		)
		return nil
	}
}
