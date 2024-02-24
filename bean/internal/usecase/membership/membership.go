package membership

import (
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

func (u *UseCase) Create(email string, createdAt time.Time) (*model.Membership, error) {
	user, err := u.Users.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by email: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	membership, err := u.Memberships.Create(user.ID, createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create membership: %v", err)
	}

	return membership, nil
}

func (u *UseCase) Cancel(userID string, expiresAt time.Time) (*model.Membership, error) {
	membership, err := u.Memberships.Update(userID, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update membership: %v", err)
	}

	if membership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	return membership, nil
}

func (u *UseCase) Validate(userID string) (bool, error) {
	if u.Bypass {
		return true, nil
	}

	membership, err := u.Memberships.Find(userID)
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
