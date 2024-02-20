package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

func createSessionCookie(sessionToken *model.SessionToken) (*http.Cookie, error) {
	json, err := sessionToken.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal session token: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(json))

	return &http.Cookie{
		Name:     "session",
		Value:    encoded,
		Expires:  sessionToken.ExpiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	}, nil
}
