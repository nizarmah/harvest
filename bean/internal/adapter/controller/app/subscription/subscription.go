package subscription

import (
	"github.com/whatis277/harvest/bean/internal/usecase/subscription"

	"github.com/whatis277/harvest/bean/internal/adapter/interfaces"
)

type Controller struct {
	Subscriptions subscription.UseCase

	CreateView interfaces.CreateSubscriptionView
	DeleteView interfaces.DeleteSubscriptionView
}
