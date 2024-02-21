package auth

import (
	"net/http"

	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	Passwordless passwordless.UseCase

	LoginView interfaces.LoginView
}

func (c *Controller) Authenticate(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := getSessionToken(r)
		if err != nil || sessionToken == nil {
			removeSessionTokenCookie(w)
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		session, err := c.Passwordless.Authenticate(sessionToken)
		if err != nil || session == nil {
			removeSessionTokenCookie(w)
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		err = addSessionTokenCookie(w, sessionToken)
		if err != nil {
			removeSessionTokenCookie(w)
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		ctx := NewContextWithSession(r.Context(), session)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (c *Controller) Authorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, _ := getSessionToken(r)
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

		err = addSessionTokenCookie(w, sessionToken)
		if err != nil {
			removeSessionTokenCookie(w)
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
