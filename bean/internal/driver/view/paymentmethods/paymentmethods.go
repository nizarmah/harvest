package paymentmethods

import (
	"harvest/bean/internal/entity"

	"harvest/bean/internal/driver/view"
)

var New = view.New[entity.PaymentMethodsViewData]
