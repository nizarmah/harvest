package handler

import (
	"fmt"
	"net/http"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/adapter/interfaces"
)

type handler struct {
	landingView interfaces.LandingView
	loginView   interfaces.LoginView
}

func New(
	landingView interfaces.LandingView,
	loginView interfaces.LoginView,
) *handler {
	return &handler{
		landingView: landingView,
		loginView:   loginView,
	}
}

func (h *handler) Landing(w http.ResponseWriter, r *http.Request) {
	e := h.landingView.Render(w, nil)
	if e != nil {
		fmt.Fprintf(w, "Error: %v", e)
		return
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	e := h.loginView.Render(w, &entity.LoginViewData{
		Email: email,
	})

	if e != nil {
		fmt.Fprintf(w, "Error: %v", e)
		return
	}
}
