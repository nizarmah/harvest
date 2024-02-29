package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"
)

func (c *Controller) SignupPage() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		sessionToken := c.getSessionToken(r)
		if sessionToken != nil {
			AuthedUserRedirect(w, r)
			return nil
		}

		err := c.SignUpView.Render(w, &viewmodel.SignUpViewData{})
		if err != nil {
			return &base.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"auth: signup: error rendering signup view: %v",
					err,
				),
			}
		}

		return nil
	}
}
