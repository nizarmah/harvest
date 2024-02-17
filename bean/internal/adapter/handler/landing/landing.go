package landing

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	view interfaces.LandingView
}

func New(view interfaces.LandingView) http.Handler {
	return &handler{
		view: view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.view.Render(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
