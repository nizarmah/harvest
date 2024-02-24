package buymeacoffee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
)

type Controller struct {
	WebhookSecret string
}

func (c *Controller) Webhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		if userAgent != "BMC-HTTPS-ROBOT" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		signature := r.Header.Get("X-Signature-Sha256")
		ok, err := c.validateSignature(body, signature)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (c *Controller) validateSignature(body []byte, signature string) (bool, error) {
	digest := hmac.New(sha256.New, []byte(c.WebhookSecret))

	_, err := digest.Write(body)
	if err != nil {
		return false, err
	}

	hex := hex.EncodeToString(digest.Sum(nil))

	return hex == signature, nil
}
