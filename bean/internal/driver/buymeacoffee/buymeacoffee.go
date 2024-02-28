package buymeacoffee

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/whatis277/harvest/bean/internal/usecase/membership"
	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"
)

type Controller struct {
	AcceptTestEvents bool
	WebhookSecret    string

	Passwordless passwordless.UseCase
	Memberships  membership.UseCase
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

		var event Event
		err = json.Unmarshal(body, &event)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !event.LiveMode && !c.AcceptTestEvents {
			w.WriteHeader(http.StatusOK)
			return
		}

		ctx := r.Context()

		switch event.Type {
		case "membership.started":
			c.membershipStarted(ctx, w, event)
			return

		case "membership.cancelled":
			c.membershipCancelled(ctx, w, event)
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

	signatureHex := hex.EncodeToString(digest.Sum(nil))

	return signatureHex == signature, nil
}

func (c *Controller) membershipStarted(
	ctx context.Context,
	w http.ResponseWriter,
	event Event,
) {
	var data MembershipStartedData
	err := json.Unmarshal(event.Data, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	email := data.SupporterEmail
	createdAt := time.Unix(data.CurrentPeriodStart, 0)

	_, err = c.Memberships.Create(ctx, email, createdAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.Passwordless.Login(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) membershipCancelled(
	ctx context.Context,
	w http.ResponseWriter,
	event Event,
) {
	var data MembershipCancelledData
	err := json.Unmarshal(event.Data, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	email := data.SupporterEmail
	expiresAt := time.Unix(data.CurrentPeriodEnd, 0)

	_, err = c.Memberships.Cancel(ctx, email, expiresAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
