package passwordless

import (
	"context"
	"fmt"
	"time"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/google/uuid"
)

type UseCase struct {
	Sender  string
	BaseURL string

	Users    interfaces.UserDataSource
	Tokens   interfaces.LoginTokenDataSource
	Sessions interfaces.SessionDataSource

	Hasher  interfaces.Hasher
	Emailer interfaces.Emailer
}

func (u *UseCase) Login(
	ctx context.Context,
	email string,
) error {
	if err := validateEmail(email); err != nil {
		return fmt.Errorf("failed to validate email: %w", err)
	}

	user, err := u.Users.FindByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if user == nil {
		return nil
	}

	token, err := u.Tokens.FindUnexpiredByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to find existing token: %w", err)
	}

	if token != nil {
		return nil
	}

	password, hash, err := u.generatePassword()
	if err != nil {
		return fmt.Errorf("failed to generate password: %w", err)
	}

	token, err = u.Tokens.Create(ctx, email, hash)
	if err != nil {
		return fmt.Errorf("failed to create token: %w", err)
	}

	if err = u.Emailer.Send(
		u.Sender,
		email,
		"Login",
		u.buildEmailBody(token.ID, password),
	); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (u *UseCase) Authorize(
	ctx context.Context,
	id string,
	password string,
) (*model.SessionToken, error) {
	token, err := u.Tokens.FindUnexpiredByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find token: %w", err)
	}

	if token == nil {
		return nil, nil
	}

	if err := u.Hasher.Compare(password, token.HashedToken); err != nil {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	}

	user, err := u.findOrCreateUser(ctx, token.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find or create user: %w", err)
	}

	csrfToken, csrfHash, err := u.generatePassword()
	if err != nil {
		return nil, fmt.Errorf("failed to generate password: %w", err)
	}

	session, err := u.Sessions.Create(ctx, user.ID, csrfHash, 14*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	sessionToken := &model.SessionToken{
		ID:        session.ID,
		Token:     csrfToken,
		ExpiresAt: session.ExpiresAt,
	}

	u.Tokens.Delete(ctx, token.ID)

	return sessionToken, nil
}

func (u *UseCase) Authenticate(
	ctx context.Context,
	sessionToken *model.SessionToken,
) (*model.Session, error) {
	session, err := u.Sessions.FindByID(ctx, sessionToken.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find session: %w", err)
	}

	if session == nil {
		return nil, nil
	}

	if err := u.Hasher.Compare(sessionToken.Token, session.HashedToken); err != nil {
		return nil, fmt.Errorf("failed to compare session token: %w", err)
	}

	err = u.Sessions.Refresh(ctx, session, 14*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh session: %w", err)
	}

	sessionToken.ExpiresAt = session.ExpiresAt

	return session, nil
}

func (u *UseCase) Logout(
	ctx context.Context,
	session *model.Session,
) error {
	err := u.Sessions.Delete(ctx, session.ID)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}

func (u *UseCase) findOrCreateUser(
	ctx context.Context,
	email string,
) (*model.User, error) {
	user, err := u.Users.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if user != nil {
		return user, nil
	}

	user, err = u.Users.Create(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (u *UseCase) buildEmailBody(tokenID, password string) string {
	authUrl := fmt.Sprintf(
		"%s/auth/%s/%s",
		u.BaseURL,
		tokenID, password,
	)

	return fmt.Sprintf(
		("Hello." +
			"\r\n\r\n" +
			"Use this link to login to Bean:" +
			"\r\n" +
			"%s" +
			"\r\n\r\n" +
			"This link will expire in 10 minutes." +
			"\r\n" +
			"If you did not request this, don't worry." +
			"\r\n\r\n" +
			"Cheers."),
		authUrl,
	)
}

func (u *UseCase) generatePassword() (string, string, error) {
	rand, err := uuid.NewRandom()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate random: %w", err)
	}

	password := rand.String()

	hash, err := u.Hasher.Hash(password)
	if err != nil {
		return "", "", fmt.Errorf("failed to hash password: %w", err)
	}

	return password, hash, nil
}
