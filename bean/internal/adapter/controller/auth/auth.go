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

func (c *Controller) Authorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, password := r.PathValue("id"), r.PathValue("password")
		if id == "" || password == "" {
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		sessionToken, err := c.Passwordless.Login(id, password)
		if err != nil {
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		cookie, err := createSessionCookie(sessionToken)
		if err != nil {
			http.Redirect(w, r, "/get-started", http.StatusSeeOther)
			return
		}

		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
