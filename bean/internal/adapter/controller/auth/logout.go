package auth

import "net/http"

func (c *Controller) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := SessionFromContext(r.Context())

		c.Passwordless.Logout(session)

		c.cleanupSessionToken(w)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
