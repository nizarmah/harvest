package auth

import (
	"fmt"
	"net/http"
)

func (c *Controller) CheckMembership(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := SessionFromContext(r.Context())

		isMember, err := c.Memberships.CheckByID(session.UserID)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		if !isMember {
			http.Redirect(w, r, "/onboarding", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}
