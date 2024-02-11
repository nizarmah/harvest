package subscription

import (
	"net/http"

	"harvest/bean/internal/usecase/subscription"

	"harvest/bean/internal/adapter/interfaces"
)

type crudHandler struct {
	Create http.Handler
	Delete http.Handler
}

func New(
	sub subscription.UseCase,
	createView interfaces.CreateSubscriptionView,
) crudHandler {
	return crudHandler{
		Create: newCreateHandler(sub, createView),
		Delete: nil,
	}
}
