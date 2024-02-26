package auth

import (
	"net/http"

	"github.com/whatis277/harvest/bean/internal/usecase/membership"
	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	BypassHTTPS bool

	Passwordless passwordless.UseCase
	Memberships  membership.UseCase

	LoginView interfaces.LoginView
}

func (c *Controller) Authenticate(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := c.getSessionToken(r)
		if err != nil || sessionToken == nil {
			c.cleanSessionToken(w)
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		session, err := c.Passwordless.Authenticate(sessionToken)
		if err != nil || session == nil {
			c.cleanSessionToken(w)
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		err = c.createSessionToken(w, sessionToken)
		if err != nil {
			c.cleanSessionToken(w)
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		ctx := NewContextWithSession(r.Context(), session)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (c *Controller) Authorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, _ := c.getSessionToken(r)
		if sessionToken != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		id, password := r.PathValue("id"), r.PathValue("password")
		if id == "" || password == "" {
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		sessionToken, err := c.Passwordless.Authorize(id, password)
		if err != nil {
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		err = c.createSessionToken(w, sessionToken)
		if err != nil {
			c.cleanSessionToken(w)
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
