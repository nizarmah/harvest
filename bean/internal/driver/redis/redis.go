package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

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

func (c *Cache) Close() {
	c.Client.Close()
}

func (c *Cache) Set(
	ctx context.Context,
	ns string,
	key string,
	value interface{},
	expiration time.Duration,
) *redis.StatusCmd {
	return c.Client.Set(ctx, prefix(ns, key), value, expiration)
}

func (c *Cache) Get(
	ctx context.Context,
	ns string,
	key string,
) *redis.StringCmd {
	return c.Client.Get(ctx, prefix(ns, key))
}

func (c *Cache) Del(
	ctx context.Context,
	ns string,
	key string,
) *redis.IntCmd {
	return c.Client.Del(ctx, prefix(ns, key))
}

func prefix(namespace, key string) string {
	return fmt.Sprintf("%s:%s", namespace, key)
}
