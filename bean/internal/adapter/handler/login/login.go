package login

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	auth passwordless.UseCase

	view interfaces.LoginView
}

func New(
	auth passwordless.UseCase,
	view interfaces.LoginView,
) http.Handler {
	return &handler{
		auth: auth,
		view: view,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.render(w, r, &viewmodel.LoginViewData{})
		return
	}

	if r.Method == http.MethodPost {
		h.post(w, r)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *handler) render(w http.ResponseWriter, r *http.Request, data *viewmodel.LoginViewData) {
	err := h.view.Render(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
}

func (h *handler) post(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email == "" {
		h.render(w, r, &viewmodel.LoginViewData{})
		return
	}

	err := h.auth.SendEmail(email)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	h.render(w, r, &viewmodel.LoginViewData{
		Email: email,
	})
}
