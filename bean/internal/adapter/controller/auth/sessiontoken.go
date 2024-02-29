package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

func (c *Controller) getSessionToken(r *http.Request) *model.SessionToken {
	cookie := c.sessionTokenCookie()
	cookie, err := r.Cookie(cookie.Name)
	if err != nil {
		return nil
	}

	decoded, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil
	}

	token := model.SessionToken{}
	err = token.UnmarshalBinary(decoded)
	if err != nil {
		return nil
	}

	return &token
}

func (c *Controller) createSessionToken(w http.ResponseWriter, token *model.SessionToken) error {
	json, err := token.MarshalBinary()
	if err != nil {
		return fmt.Errorf("create session token: marshal binary: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(json))

	cookie := c.sessionTokenCookie()

	cookie.Value = encoded
	cookie.Expires = token.ExpiresAt

	http.SetCookie(w, cookie)

	return nil
}

func (c *Controller) cleanupSessionToken(w http.ResponseWriter) {
	cookie := c.sessionTokenCookie()
	cookie.MaxAge = -1

	http.SetCookie(w, cookie)
}

func (c *Controller) sessionTokenCookie() *http.Cookie {
	if c.BypassHTTPS {
		return &http.Cookie{
			Name:     "session",
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		}
	}

	return &http.Cookie{
		Name:     "__Host-session",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}
