package auth

import (
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

var (
	defaultObscureStatus = http.StatusFound

	defaultAuthedRedirectPath   = "/home"
	defaultUnauthedRedirectPath = "/logout"
)

func NewUnauthorizedError(msg string) error {
	return &model.HTTPError{
		Status:       defaultObscureStatus,
		RedirectPath: defaultUnauthedRedirectPath,
		Message:      msg,
	}
}

func NewAuthorizedError(msg string) error {
	return &model.HTTPError{
		Status:       defaultObscureStatus,
		RedirectPath: defaultAuthedRedirectPath,
		Message:      msg,
	}
}
