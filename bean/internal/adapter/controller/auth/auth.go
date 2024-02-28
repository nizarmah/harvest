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

	LoginView  interfaces.LoginView
	SignUpView interfaces.SignUpView
}

func (c *Controller) Authenticate(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := c.getSessionToken(r)
		if err != nil || sessionToken == nil {
			c.cleanupSessionToken(w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx := r.Context()

		session, err := c.Passwordless.Authenticate(ctx, sessionToken)
		if err != nil || session == nil {
			c.cleanupSessionToken(w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		err = c.createSessionToken(w, sessionToken)
		if err != nil {
			c.cleanupSessionToken(w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		sessionCtx := NewContextWithSession(ctx, session)

		next.ServeHTTP(w, r.WithContext(sessionCtx))
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
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx := r.Context()

		sessionToken, err := c.Passwordless.Authorize(ctx, id, password)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		err = c.createSessionToken(w, sessionToken)
		if err != nil {
			c.cleanupSessionToken(w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
