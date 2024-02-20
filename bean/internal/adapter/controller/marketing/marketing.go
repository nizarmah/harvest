package marketing

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	LandingView interfaces.LandingView
}

func (c *Controller) LandingPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := c.LandingView.Render(w, &viewmodel.LandingViewData{})
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
	}
}
