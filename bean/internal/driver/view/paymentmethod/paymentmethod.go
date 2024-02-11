package login

import (
	"harvest/bean/internal/entity/viewmodel"

	"harvest/bean/internal/driver/view"
)

var NewCreate = view.New[viewmodel.CreatePaymentMethodViewData]
var NewDelete = view.New[viewmodel.DeletePaymentMethodViewData]
