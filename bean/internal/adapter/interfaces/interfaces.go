package interfaces

import (
	"net/http"

	"harvest/bean/internal/entity"
)

// --- Views ---

type View[T any] interface {
	Render(http.ResponseWriter, *T) error
}

type LandingView View[entity.LandingViewData]
type LoginView View[entity.LoginViewData]
