package marketing

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	LandingView interfaces.LandingView
}

func (c *Controller) LandingPage() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := c.LandingView.Render(w, &viewmodel.LandingViewData{})
		if err != nil {
			return &base.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"landing: error rendering landing page: %v",
					err,
				),
			}
		}

		return nil
	}
}
