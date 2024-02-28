package membership

import (
	"context"
	"fmt"
	"time"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"
)

type UseCase struct {
	Bypass bool

	Users       interfaces.UserDataSource
	Memberships interfaces.MembershipDataSource
}

func (u *UseCase) Create(
	ctx context.Context,
	email string,
	createdAt time.Time,
) (*model.Membership, error) {
	user, err := u.findOrCreateUser(email)
	if err != nil {
		return nil, fmt.Errorf("failed to find or create user: %v", err)
	}

	membership, err := u.Memberships.Create(ctx, user.ID, createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create membership: %v", err)
	}

	return membership, nil
}

func (u *UseCase) Cancel(
	ctx context.Context,
	email string,
	expiresAt time.Time,
) (*model.Membership, error) {
	user, err := u.Users.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	membership, err := u.Memberships.Update(ctx, user.ID, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update membership: %v", err)
	}

	if membership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	return membership, nil
}

func (u *UseCase) CheckByID(
	ctx context.Context,
	userID string,
) (bool, error) {
	if u.Bypass {
		return true, nil
	}

	membership, err := u.Memberships.Find(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("failed to find membership: %v", err)
	}

	if membership == nil {
		return false, nil
	}

	if membership.ExpiresAt != nil && membership.ExpiresAt.Before(time.Now()) {
		return false, nil
	}

	return true, nil
}

func (u *UseCase) CheckByEmail(
	ctx context.Context,
	email string,
) (bool, error) {
	if u.Bypass {
		return true, nil
	}

	user, err := u.Users.FindByEmail(email)
	if err != nil {
		return false, fmt.Errorf("failed to find user: %v", err)
	}

	if user == nil {
		return false, nil
	}

	return u.CheckByID(ctx, user.ID)
}

func (u *UseCase) findOrCreateUser(email string) (*model.User, error) {
	user, err := u.Users.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if user != nil {
		return user, nil
	}

	user, err = u.Users.Create(email)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
