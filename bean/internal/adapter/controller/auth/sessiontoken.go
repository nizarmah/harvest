package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

const sessionCookieName = "session"

func getSessionToken(r *http.Request) (*model.SessionToken, error) {
	cookie, err := r.Cookie(sessionCookieName)
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

func addSessionTokenCookie(w http.ResponseWriter, token *model.SessionToken) error {
	json, err := token.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal session token: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(json))

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    encoded,
		Expires:  token.ExpiresAt,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return nil
}

func removeSessionTokenCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}
