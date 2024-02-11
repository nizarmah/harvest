package paymentmethod

import (
	"net/http"

	estimatorUC "harvest/bean/internal/usecase/estimator"
	"harvest/bean/internal/usecase/paymentmethod"

	"harvest/bean/internal/adapter/interfaces"
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
