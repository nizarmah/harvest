package auth

import "net/http"

func (c *Controller) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := SessionFromContext(r.Context())

		c.Passwordless.Logout(session)

		removeSessionTokenCookie(w)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
