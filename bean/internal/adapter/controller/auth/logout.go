package auth

import (
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
		if session != nil {
			c.Passwordless.Logout(ctx, session)
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
