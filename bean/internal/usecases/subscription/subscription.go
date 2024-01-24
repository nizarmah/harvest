package subscription

import (
	"fmt"
	"harvest/bean/internal/entity"
	"harvest/bean/internal/usecases/interfaces"
)

type UseCase struct {
	paymentMethods interfaces.PaymentMethodDataSource
	subscriptions  interfaces.SubscriptionDataSource
}

func (u *UseCase) Create(
	userID string,
	paymentMethodID string,
	label string,
	provider string,
	amount int,
	interval int,
	period string,
) (*entity.Subscription, error) {
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

	method, err := u.paymentMethods.FindByID(userID, paymentMethodID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment method: %w", err)
	}

	if method == nil {
		return nil, fmt.Errorf("payment method not found")
	}

	subscription, err := u.subscriptions.Create(
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

func (u *UseCase) Get(userID string, subscriptionID string) (*entity.Subscription, error) {
	subscription, err := u.subscriptions.FindByID(userID, subscriptionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}

	return subscription, nil
}

func (u *UseCase) List(userID string) ([]*entity.Subscription, error) {
	subscriptions, err := u.subscriptions.FindByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscriptions: %w", err)
	}

	return subscriptions, nil
}

func (u *UseCase) Delete(userID string, subscriptionID string) error {
	if err := u.subscriptions.Delete(userID, subscriptionID); err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}

	return nil
}
