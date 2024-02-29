package auth

import (
	"net/http"
)

var (
	defaultObscureStatus = http.StatusFound

	defaultAuthedRedirectPath   = "/home"
	defaultUnauthedRedirectPath = "/logout"
)

func UnauthedUserRedirect(
	w http.ResponseWriter,
	r *http.Request,
) {
	http.Redirect(
		w,
		r,
		defaultUnauthedRedirectPath,
		defaultObscureStatus,
	)
}

func AuthedUserRedirect(
	w http.ResponseWriter,
	r *http.Request,
) {
	http.Redirect(
		w,
		r,
		defaultAuthedRedirectPath,
		defaultObscureStatus,
	)
}
