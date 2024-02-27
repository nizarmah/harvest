package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"
)

func (c *Controller) SignupPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionToken, _ := c.getSessionToken(r)
		if sessionToken != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		err := c.SignUpView.Render(w, &viewmodel.SignUpViewData{})
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
		}
	}
}
