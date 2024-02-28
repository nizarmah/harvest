package subscription

import (
	"context"
	"fmt"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"
)

type UseCase struct {
	PaymentMethods interfaces.PaymentMethodDataSource
	Subscriptions  interfaces.SubscriptionDataSource
}

func (u *UseCase) Create(
	ctx context.Context,
	userID string,
	paymentMethodID string,
	label string,
	provider string,
	amount int,
	interval int,
	period model.SubscriptionPeriod,
) (*model.Subscription, error) {
	if err := validateLabel(label); err != nil {
		return nil, fmt.Errorf("invalid label: %w", err)
	}

	if err := validateProvider(provider); err != nil {
		return nil, fmt.Errorf("invalid provider: %w", err)
	}

	if err := validateAmount(amount); err != nil {
		return nil, fmt.Errorf("invalid amount: %w", err)
	}

	if err := validateInterval(interval); err != nil {
		return nil, fmt.Errorf("invalid interval: %w", err)
	}

	if err := validatePeriod(period); err != nil {
		return nil, fmt.Errorf("invalid period: %w", err)
	}

	method, err := u.PaymentMethods.FindByID(ctx, userID, paymentMethodID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment method: %w", err)
	}

	if method == nil {
		return nil, fmt.Errorf("payment method not found")
	}

	subscription, err := u.Subscriptions.Create(
		ctx,
		userID,
		paymentMethodID,
		label,
		provider,
		amount,
		interval,
		period,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}

	return subscription, nil
}

func (u *UseCase) Get(
	ctx context.Context,
	userID string,
	subscriptionID string,
) (*model.Subscription, error) {
	subscription, err := u.Subscriptions.FindByID(ctx, userID, subscriptionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	return subscription, nil
}

func (u *UseCase) Delete(
	ctx context.Context,
	userID string,
	subscriptionID string,
) error {
	if err := u.Subscriptions.Delete(ctx, userID, subscriptionID); err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}

	return nil
}
