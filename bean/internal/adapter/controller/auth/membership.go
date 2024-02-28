package auth

import (
	"fmt"
	"net/http"
)

func (c *Controller) CheckMembership(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		session := SessionFromContext(ctx)
		if session == nil {
			http.Redirect(w, r, "/logout", http.StatusFound)
			return
		}

		isMember, err := c.Memberships.CheckByID(ctx, session.UserID)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		if !isMember {
			http.Redirect(w, r, "/renew-plan", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}
