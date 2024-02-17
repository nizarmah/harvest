package subscription

import (
	"github.com/whatis277/harvest/bean/internal/entity/viewmodel"

	"github.com/whatis277/harvest/bean/internal/driver/view"
)

var NewCreate = view.New[viewmodel.CreateSubscriptionViewData]
var NewDelete = view.New[viewmodel.DeleteSubscriptionViewData]
