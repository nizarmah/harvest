package auth

import (
	"net/http"

	"harvest/bean/internal/usecase/passwordless"
)

type handler struct {
	auth passwordless.UseCase
}

func New(
	auth passwordless.UseCase,
) http.Handler {
	return &handler{
		auth: auth,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.get(w, r)
		return
	}

	http.Redirect(w, r, "/get-started", http.StatusSeeOther)
}

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	id, password := r.FormValue("i"), r.FormValue("p")
	if id == "" || password == "" {
		http.Redirect(w, r, "/get-started", http.StatusSeeOther)
		return
	}

	_, err := h.auth.Login(id, password)
	if err != nil {
		http.Redirect(w, r, "/get-started", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
