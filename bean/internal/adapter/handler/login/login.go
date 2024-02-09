package landing

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity/viewmodel"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	view interfaces.LoginView
}

func New(view interfaces.LoginView) http.Handler {
	return &handler{
		view: view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	err := h.view.Render(w, &viewmodel.LoginViewData{
		Email: email,
	})
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
