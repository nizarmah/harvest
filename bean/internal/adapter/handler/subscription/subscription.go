package subscription

import (
	"net/http"

	"github.com/whatis277/harvest/bean/internal/usecase/subscription"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type crudHandler struct {
	Create http.Handler
	Delete http.Handler
}

func New(
	sub subscription.UseCase,
	createView interfaces.CreateSubscriptionView,
	deleteView interfaces.DeleteSubscriptionView,
) crudHandler {
	return crudHandler{
		Create: newCreateHandler(sub, createView),
		Delete: newDeleteHandler(sub, deleteView),
	}
}
