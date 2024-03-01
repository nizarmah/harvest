package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/usecase/membership"
	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	BypassHTTPS bool

	Passwordless passwordless.UseCase
	Memberships  membership.UseCase

	LoginView  interfaces.LoginView
	SignUpView interfaces.SignUpView
}

func (c *Controller) Authenticate(next base.HTTPHandler) base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		sessionToken := c.getSessionToken(r)
		if sessionToken == nil {
			UnauthedUserRedirect(w, r)
			return nil
		}

		ctx := r.Context()

		session, err := c.Passwordless.Authenticate(ctx, sessionToken)
		if err != nil {
			UnauthedUserRedirect(w, r)
			return &base.HTTPError{
				Message: fmt.Sprintf(
					"auth: authenticate: error authenticating user: %v",
					err,
				),
			}
		}

		err = c.createSessionToken(w, sessionToken)
		if err != nil {
			UnauthedUserRedirect(w, r)
			return &base.HTTPError{
				Message: fmt.Sprintf(
					"auth: authenticate: error creating session token: %v",
					err,
				),
			}
		}

		sessionCtx := NewContextWithSession(ctx, session)

		return next(w, r.WithContext(sessionCtx))
	}
}

func (c *Controller) Authorize() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		sessionToken := c.getSessionToken(r)
		if sessionToken != nil {
			AuthedUserRedirect(w, r)
			return nil
		}

		id, password := r.PathValue("id"), r.PathValue("password")
		if id == "" || password == "" {
			UnauthedUserRedirect(w, r)
			return nil
		}

		ctx := r.Context()

		sessionToken, err := c.Passwordless.Authorize(ctx, id, password)
		if err != nil {
			UnauthedUserRedirect(w, r)
			return &base.HTTPError{
				Message: fmt.Sprintf(
					"auth: authorize: error authorizing user: %v",
					err,
				),
			}
		}

		err = c.createSessionToken(w, sessionToken)
		if err != nil {
			UnauthedUserRedirect(w, r)
			return &base.HTTPError{
				Message: fmt.Sprintf(
					"auth: authorize: error creating session token: %v",
					err,
				),
			}
		}

		AuthedUserRedirect(w, r)
		return nil
	}
}
