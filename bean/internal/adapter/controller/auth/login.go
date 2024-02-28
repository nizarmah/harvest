package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"
)

func (c *Controller) LoginPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, _ := c.getSessionToken(r)
		if sessionToken != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		c.renderLogin(w, &viewmodel.LoginViewData{})
	}
}

func (c *Controller) LoginForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, _ := c.getSessionToken(r)
		if sessionToken != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		email := r.FormValue("email")
		if email == "" {
			c.LoginView.Render(w, &viewmodel.LoginViewData{})
			return
		}

		password := r.FormValue("password")
		if password != "" {
			c.LoginView.Render(w, &viewmodel.LoginViewData{})
			return
		}

		ctx := r.Context()

		err := c.Passwordless.Login(ctx, email)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		c.renderLogin(w, &viewmodel.LoginViewData{
			Email: email,
		})
	}
}

func (c *Controller) renderLogin(w http.ResponseWriter, data *viewmodel.LoginViewData) {
	err := c.LoginView.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
}
