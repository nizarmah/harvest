package token

import (
	"errors"

	"harvest/bean/internal/entity"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/database"
)

type dataSource struct {
	db *database.DB
}

func New(db *database.DB) usecase.LoginTokenDataSource {
	return &dataSource{
		db: db,
	}
}

func (ds *dataSource) Create(inputToken *entity.LoginToken) error {
	if len(inputToken.HashedToken) > 60 {
		return errors.New("token is too long")
	}

	_, err := ds.db.Pool.Exec(
		("INSERT INTO login_tokens (email, hashed_token, expires_at) " +
			"VALUES (?, ?, DATE_ADD(NOW(), INTERVAL 10 MINUTE)) " +
			"ON DUPLICATE KEY UPDATE hashed_token = ?, created_at = NOW(), expires_at = DATE_ADD(NOW(), INTERVAL 10 MINUTE)"),
		inputToken.Email, inputToken.HashedToken, inputToken.HashedToken,
	)
	if err != nil {
		return errors.New("error creating login token")
	}

	return nil
}

func (ds *dataSource) FindUnexpired(token *entity.LoginToken) (*entity.LoginToken, error) {
	t := &entity.LoginToken{}

	err := ds.db.Pool.
		QueryRow(
			"SELECT * FROM login_tokens WHERE email = ? AND BINARY hashed_token = ? AND expires_at > NOW()",
			token.Email, token.HashedToken,
		).
		Scan(&t.ID, &t.Email, &t.HashedToken, &t.CreatedAt, &t.ExpiresAt)

	if err != nil {
		return nil, errors.New("error finding unexpired login token")
	}

	return t, nil
}

func (ds *dataSource) Delete(token *entity.LoginToken) error {
	_, err := ds.db.Pool.Exec("DELETE FROM login_tokens WHERE id = ?", token.ID)
	if err != nil {
		return errors.New("error deleting login token")
	}

	return nil
}
