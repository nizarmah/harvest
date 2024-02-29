package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// --- HTTP ---

// --- HTTP --- Error ---

type HTTPError struct {
	Status       int    `json:"status"`
	RedirectPath string `json:"redirect_path"`

	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf(
			"error marshalling http error: %v",
			err,
		)
	}

	return string(b)
}

// --- HTTP --- Handler ---

type HTTPHandler = func(w http.ResponseWriter, r *http.Request) error
