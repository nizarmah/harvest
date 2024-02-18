package auth

import (
	"encoding/base64"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"
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

	sessionToken, err := h.auth.Login(id, password)
	if err != nil {
		http.Redirect(w, r, "/get-started", http.StatusSeeOther)
		return
	}

	sessionBin, err := sessionToken.MarshalBinary()
	if err != nil {
		http.Redirect(w, r, "/get-started", http.StatusSeeOther)
		return
	}

	sessionVal := base64.StdEncoding.EncodeToString([]byte(sessionBin))
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionVal,
		Expires:  sessionToken.ExpiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
