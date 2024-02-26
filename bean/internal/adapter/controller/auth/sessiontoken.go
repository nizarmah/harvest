package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

func (c *Controller) getSessionToken(r *http.Request) (*model.SessionToken, error) {
	cookie := c.sessionTokenCookie()
	cookie, err := r.Cookie(cookie.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get session token: %w", err)
	}

	decoded, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to decode session token: %w", err)
	}

	token := model.SessionToken{}
	err = token.UnmarshalBinary(decoded)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal session token: %w", err)
	}

	return &token, nil
}

func (c *Controller) createSessionToken(w http.ResponseWriter, token *model.SessionToken) error {
	json, err := token.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal session token: %w", err)
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
