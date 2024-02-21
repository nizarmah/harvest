package passwordless

import (
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

func (u *UseCase) SendEmail(email string) error {
	if err := validateEmail(email); err != nil {
		return fmt.Errorf("failed to validate email: %w", err)
	}

	rand, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate password: %w", err)
	}

	password := rand.String()
	hash, err := u.Hasher.Hash(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	token, err := u.Tokens.Create(email, hash)
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

func (u *UseCase) Login(id string, password string) (*model.SessionToken, error) {
	token, err := u.Tokens.FindUnexpired(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find token: %w", err)
	}

	if token == nil {
		return nil, fmt.Errorf("token not found")
	}

	if err := u.Hasher.Compare(password, token.HashedToken); err != nil {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	}

	user, err := u.findOrCreateUser(token.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find or create user: %w", err)
	}

	session, err := u.Sessions.Create(user.ID, token.HashedToken, 14*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	sessionToken := &model.SessionToken{
		ID:        session.ID,
		Token:     password,
		ExpiresAt: session.ExpiresAt,
	}

	u.Tokens.Delete(token.ID)

	return sessionToken, nil
}

func (u *UseCase) Authenticate(sessionToken *model.SessionToken) (*model.Session, error) {
	session, err := u.Sessions.FindByID(sessionToken.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find session: %w", err)
	}

	if session == nil {
		return nil, fmt.Errorf("session not found")
	}

	if err := u.Hasher.Compare(sessionToken.Token, session.HashedToken); err != nil {
		return nil, fmt.Errorf("failed to compare session token: %w", err)
	}

	err = u.Sessions.Refresh(session, 14*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh session: %w", err)
	}

	sessionToken.ExpiresAt = session.ExpiresAt

	return session, nil
}

func (u *UseCase) Logout(sessionToken *model.SessionToken) error {
	session, err := u.Sessions.FindByID(sessionToken.ID)
	if err != nil {
		return fmt.Errorf("failed to find session: %w", err)
	}

	if session == nil {
		return fmt.Errorf("session not found")
	}

	if err := u.Hasher.Compare(sessionToken.Token, session.HashedToken); err != nil {
		return fmt.Errorf("failed to compare session token: %w", err)
	}

	u.Sessions.Delete(session.ID)

	return nil
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
