package paymentmethod

import (
	"github.com/whatis277/harvest/bean/internal/usecase/estimator"
	"github.com/whatis277/harvest/bean/internal/usecase/paymentmethod"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	Estimator      estimator.UseCase
	PaymentMethods paymentmethod.UseCase

	CreateView interfaces.CreatePaymentMethodView
	DeleteView interfaces.DeletePaymentMethodView
}
