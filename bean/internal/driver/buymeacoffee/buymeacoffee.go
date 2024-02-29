package buymeacoffee

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/whatis277/harvest/bean/internal/usecase/membership"
	"github.com/whatis277/harvest/bean/internal/usecase/passwordless"

	"github.com/whatis277/harvest/bean/internal/adapter/controller/base"
)

type Controller struct {
	AcceptTestEvents bool
	WebhookSecret    string

	Passwordless passwordless.UseCase
	Memberships  membership.UseCase
}

func (c *Controller) Webhook() base.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		userAgent := r.Header.Get("User-Agent")
		if userAgent != "BMC-HTTPS-ROBOT" {
			return &base.HTTPError{
				Status: http.StatusUnauthorized,

				Message: "buymeacoffee: webhook: invalid user agent",
			}
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return &base.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"buymeacoffee: webhook: error reading body: %v",
					err,
				),
			}
		}

		signature := r.Header.Get("X-Signature-Sha256")
		ok, err := c.validateSignature(body, signature)
		if err != nil {
			return &base.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"buymeacoffee: webhook: error validating signature: %v",
					err,
				),
			}
		}

		if !ok {
			return &base.HTTPError{
				Status: http.StatusUnauthorized,

				Message: "buymeacoffee: webhook: invalid signature",
			}
		}

		var event Event
		err = json.Unmarshal(body, &event)
		if err != nil {
			return &base.HTTPError{
				Status: http.StatusInternalServerError,

				Message: fmt.Sprintf(
					"buymeacoffee: webhook: error unmarshalling event: %v",
					err,
				),
			}
		}

		if !event.LiveMode && !c.AcceptTestEvents {
			return &base.HTTPError{
				Status: http.StatusOK,

				Message: "buymeacoffee: webhook: test event not allowed",
			}
		}

		ctx := r.Context()

		switch event.Type {
		case "membership.started":
			return c.membershipStarted(ctx, w, event)

		case "membership.cancelled":
			return c.membershipCancelled(ctx, w, event)
		}

		w.WriteHeader(http.StatusOK)

		return nil
	}
}

func (c *Controller) validateSignature(
	body []byte,
	signature string,
) (bool, error) {
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
) error {
	var data MembershipStartedData
	err := json.Unmarshal(event.Data, &data)
	if err != nil {
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"buymeacoffee: webhook: error unmarshalling membership started data: %v",
				err,
			),
		}
	}

	email := data.SupporterEmail
	createdAt := time.Unix(data.CurrentPeriodStart, 0)

	_, err = c.Memberships.Create(ctx, email, createdAt)
	if err != nil {
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"buymeacoffee: webhook: error creating membership: %v",
				err,
			),
		}
	}

	err = c.Passwordless.Login(ctx, email)
	if err != nil {
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"buymeacoffee: webhook: error logging in user: %v",
				err,
			),
		}
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (c *Controller) membershipCancelled(
	ctx context.Context,
	w http.ResponseWriter,
	event Event,
) error {
	var data MembershipCancelledData
	err := json.Unmarshal(event.Data, &data)
	if err != nil {
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"buymeacoffee: webhook: error unmarshalling membership cancelled data: %v",
				err,
			),
		}
	}

	email := data.SupporterEmail
	expiresAt := time.Unix(data.CurrentPeriodEnd, 0)

	_, err = c.Memberships.Cancel(ctx, email, expiresAt)
	if err != nil {
		return &base.HTTPError{
			Status: http.StatusInternalServerError,

			Message: fmt.Sprintf(
				"buymeacoffee: webhook: error cancelling membership: %v",
				err,
			),
		}
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
