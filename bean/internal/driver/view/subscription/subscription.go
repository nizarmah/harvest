package subscription

import (
	"harvest/bean/internal/entity/viewmodel"

	"harvest/bean/internal/driver/view"
)

var NewCreate = view.New[viewmodel.CreateSubscriptionViewData]
var NewDelete = view.New[viewmodel.DeleteSubscriptionViewData]
