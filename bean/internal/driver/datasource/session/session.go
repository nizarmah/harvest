package session

import (
	"context"
	"fmt"
	"time"

	"github.com/whatis277/harvest/bean/internal/entity/model"

	"github.com/whatis277/harvest/bean/internal/usecase/interfaces"

	"github.com/whatis277/harvest/bean/internal/driver/redis"

	"github.com/google/uuid"
)

type dataSource struct {
	cache *redis.Cache
}

func New(cache *redis.Cache) interfaces.SessionDataSource {
	return &dataSource{
		cache: cache,
	}
}

func (ds *dataSource) Create(
	ctx context.Context,
	userID string,
	hashedToken string,
	duration time.Duration,
) (*model.Session, error) {
	return ds.doCreate(ctx, userID, hashedToken, duration, 0)
}

func (ds *dataSource) doCreate(
	ctx context.Context,
	userID string,
	hashedToken string,
	duration time.Duration,
	attempts int,
) (*model.Session, error) {
	if attempts > 3 {
		return nil, fmt.Errorf("failed to generate session: max attempts exceeded")
	}

	rand, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate random session id: %w", err)
	}

	id := rand.String()

	existing, err := ds.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find existing session: %w", err)
	}

	if existing != nil {
		return ds.doCreate(ctx, userID, hashedToken, duration, attempts+1)
	}

	createdAt := time.Now()
	updatedAt := createdAt
	expiresAt := updatedAt.Add(duration)

	session := &model.Session{
		ID:          id,
		UserID:      userID,
		HashedToken: hashedToken,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		ExpiresAt:   expiresAt,
	}

	err = ds.cache.Client.Set(ctx, id, session, duration).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to set session in cache: %w", err)
	}

	return session, nil
}

func (ds *dataSource) FindByID(
	ctx context.Context,
	id string,
) (*model.Session, error) {
	session := &model.Session{}

	err := ds.cache.Client.Get(ctx, id).Scan(session)
	if err == redis.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get session from cache: %w", err)
	}

	return session, nil
}

func (ds *dataSource) Refresh(
	ctx context.Context,
	session *model.Session,
	duration time.Duration,
) error {
	session.UpdatedAt = time.Now()
	session.ExpiresAt = session.UpdatedAt.Add(duration)

	err := ds.cache.Client.Set(ctx, session.ID, session, duration).Err()
	if err != nil {
		return fmt.Errorf("failed to refresh session in cache: %w", err)
	}

	return nil
}

func (ds *dataSource) Delete(
	ctx context.Context,
	id string,
) error {
	err := ds.cache.Client.Del(ctx, id).Err()
	if err != nil {
		return fmt.Errorf("failed to delete session from cache: %w", err)
	}

	return nil
}
