package passwordless

import (
	"fmt"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecases/interfaces"

	"github.com/google/uuid"
)

type UseCase struct {
	users  interfaces.UserDataSource
	tokens interfaces.LoginTokenDataSource

	hasher  interfaces.Hasher
	emailer interfaces.Emailer
}

func (u *UseCase) SendEmail(email string) error {
	rand, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate password: %w", err)
	}

	password := rand.String()
	hash, err := u.hasher.Hash(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	token, err := u.tokens.Create(email, hash)
	if err != nil {
		return fmt.Errorf("failed to create token: %w", err)
	}

	if err = u.emailer.Send(
		email,
		"Login to Bean",
		fmt.Sprintf(
			("Use the following link to login to Bean:\n"+
				"https://localhost:8080/login?i=%s&p=%s"),
			token.ID, password,
		),
	); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (u *UseCase) Login(id string, password string) (*entity.User, error) {
	hash, err := u.hasher.Hash(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	token, err := u.tokens.FindUnexpired(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find token: %w", err)
	}

	if err := u.hasher.Compare(hash, token.HashedToken); err != nil {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	}

	user, err := u.findOrCreateUser(token.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find or create user: %w", err)
	}

	u.tokens.Delete(token.ID)

	return user, nil
}

func (u *UseCase) findOrCreateUser(email string) (*entity.User, error) {
	user, err := u.users.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if user != nil {
		return user, nil
	}

	user, err = u.users.Create(email)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
