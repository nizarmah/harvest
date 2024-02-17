package paymentmethod

import (
	"net/http"

	estimatorUC "github.com/whatis277/harvest/bean/internal/usecase/estimator"
	"github.com/whatis277/harvest/bean/internal/usecase/paymentmethod"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type crudHandler struct {
	Create http.Handler
	Delete http.Handler
}

func New(
	es estimatorUC.UseCase,
	pm paymentmethod.UseCase,
	createView interfaces.CreatePaymentMethodView,
	deleteView interfaces.DeletePaymentMethodView,
) crudHandler {
	return crudHandler{
		Create: newCreateHandler(pm, createView),
		Delete: newDeleteHandler(es, pm, deleteView),
	}
}
