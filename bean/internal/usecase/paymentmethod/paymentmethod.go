package paymentmethod

import (
	"context"
	"fmt"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"
)

type UseCase struct {
	PaymentMethods interfaces.PaymentMethodDataSource
}

func (u *UseCase) Create(
	ctx context.Context,
	userID string,
	label string,
	last4 string,
	brand model.PaymentMethodBrand,
	expMonth int,
	expYear int,
) (*model.PaymentMethod, error) {
	if err := validateLabel(label); err != nil {
		return nil, fmt.Errorf("invalid label: %w", err)
	}

	if err := validateLast4(last4); err != nil {
		return nil, fmt.Errorf("invalid last4: %w", err)
	}

	if err := validateBrand(brand); err != nil {
		return nil, fmt.Errorf("invalid brand: %w", err)
	}

	if err := validateExpMonth(expMonth); err != nil {
		return nil, fmt.Errorf("invalid exp month: %w", err)
	}

	if err := validateExpYear(expYear); err != nil {
		return nil, fmt.Errorf("invalid exp year: %w", err)
	}

	method, err := u.PaymentMethods.Create(ctx, userID, label, last4, brand, expMonth, expYear)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment method: %w", err)
	}

	return method, nil
}

func (u *UseCase) Get(
	ctx context.Context,
	userID string,
	paymentMethodID string,
) (*model.PaymentMethodWithSubscriptions, error) {
	method, err := u.PaymentMethods.FindByID(ctx, userID, paymentMethodID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment method: %w", err)
	}

	return method, nil
}

func (u *UseCase) List(
	ctx context.Context,
	userID string,
) ([]*model.PaymentMethodWithSubscriptions, error) {
	methods, err := u.PaymentMethods.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list payment methods: %w", err)
	}

	return methods, nil
}

func (u *UseCase) Delete(
	ctx context.Context,
	userID string,
	paymentMethodID string,
) error {
	if err := u.PaymentMethods.Delete(ctx, userID, paymentMethodID); err != nil {
		return fmt.Errorf("failed to delete payment method: %w", err)
	}

	return nil
}
