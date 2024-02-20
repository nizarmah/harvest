package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"
)

func (c *Controller) LoginPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.renderLogin(w, &viewmodel.LoginViewData{})
	}
}

func (c *Controller) LoginForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		if email == "" {
			c.LoginView.Render(w, &viewmodel.LoginViewData{})
			return
		}

		err := c.Passwordless.SendEmail(email)
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
