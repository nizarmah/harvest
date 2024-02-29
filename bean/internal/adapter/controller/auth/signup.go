package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"
)

func (c *Controller) SignupPage() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		sessionToken := c.getSessionToken(r)
		if sessionToken != nil {
			return NewAuthorizedError(
				"auth: signup: user already has a session token",
			)
		}

		err := c.SignUpView.Render(w, &viewmodel.SignUpViewData{})
		if err != nil {
			return &model.HTTPError{
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
