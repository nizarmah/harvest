package auth

import (
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/membership"
	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	BypassHTTPS bool

	Passwordless passwordless.UseCase
	Memberships  membership.UseCase

	LoginView  interfaces.LoginView
	SignUpView interfaces.SignUpView
}

func (c *Controller) Authenticate(next model.HTTPHandler) model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		sessionToken := c.getSessionToken(r)
		if sessionToken == nil {
			return NewUnauthorizedError(
				"auth: authenticate: user has no session token",
			)
		}

		ctx := r.Context()

		session, err := c.Passwordless.Authenticate(ctx, sessionToken)
		if err != nil {
			// FIXME: This should check for a specific error type
			return NewUnauthorizedError(
				fmt.Sprintf(
					"auth: authenticate: error authenticating user: %v",
					err,
				),
			)
		}

		if session == nil {
			return NewUnauthorizedError(
				"auth: authenticate: user has no session",
			)
		}

		err = c.createSessionToken(w, sessionToken)
		if err != nil {
			NewUnauthorizedError(
				fmt.Sprintf(
					"auth: authenticate: error creating session token: %v",
					err,
				),
			)
		}

		sessionCtx := NewContextWithSession(ctx, session)

		return next(w, r.WithContext(sessionCtx))
	}
}

func (c *Controller) Authorize() model.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		sessionToken := c.getSessionToken(r)
		if sessionToken != nil {
			return NewAuthorizedError(
				"auth: authorize: user has a session token",
			)
		}

		id, password := r.PathValue("id"), r.PathValue("password")
		if id == "" || password == "" {
			return NewUnauthorizedError(
				"auth: authorize: id and password are missing",
			)
		}

		ctx := r.Context()

		sessionToken, err := c.Passwordless.Authorize(ctx, id, password)
		if err != nil {
			// FIXME: This should check for a specific error type
			return NewUnauthorizedError(
				fmt.Sprintf(
					"auth: authorize: error authorizing user: %v",
					err,
				),
			)
		}

		err = c.createSessionToken(w, sessionToken)
		if err != nil {
			return NewUnauthorizedError(
				fmt.Sprintf(
					"auth: authorize: error creating session token: %v",
					err,
				),
			)
		}

		http.Redirect(
			w,
			r,
			defaultAuthedRedirectPath,
			defaultObscureStatus,
		)

		return nil
	}
}
