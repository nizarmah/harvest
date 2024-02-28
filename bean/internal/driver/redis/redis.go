package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Cache struct {
	Client *redis.Client
}

func New(ctx context.Context, cfg *Config) (*Cache, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	client := redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Username: cfg.Username,
			Password: cfg.Password,
		},
	)
	if client == nil {
		return nil, errors.New("error creating redis client")
	}

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, errors.New("error pinging redis: " + err.Error())
	}

	return &Cache{
		Client: client,
	}, nil
}

func (cache *Cache) Close() {
	cache.Client.Close()
}
