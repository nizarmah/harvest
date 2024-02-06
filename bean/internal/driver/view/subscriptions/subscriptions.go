package subscriptions

import (
	"harvest/bean/internal/entity"

	"harvest/bean/internal/driver/view"
)

var New = view.New[entity.SubscriptionsViewData]
