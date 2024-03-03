package redis

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	TLSDisabled bool
}

type Cache struct {
	client *redis.Client
}

func New(ctx context.Context, cfg *Config) (*Cache, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	opts := &redis.Options{
		Addr:             addr,
		Username:         cfg.Username,
		Password:         cfg.Password,
		DisableIndentity: true,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}
	if cfg.TLSDisabled {
		opts.TLSConfig = nil
	}

	client := redis.NewClient(opts)
	if client == nil {
		return nil, errors.New("error creating redis client")
	}

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, errors.New("error pinging redis: " + err.Error())
	}

	return &Cache{
		client: client,
	}, nil
}

func (c *Cache) Close() {
	c.client.Close()
}

func (c *Cache) Set(
	ctx context.Context,
	ns string,
	key string,
	value interface{},
	expiration time.Duration,
) *redis.StatusCmd {
	return c.client.Set(ctx, prefix(ns, key), value, expiration)
}

func (c *Cache) Get(
	ctx context.Context,
	ns string,
	key string,
) *redis.StringCmd {
	return c.client.Get(ctx, prefix(ns, key))
}

func (c *Cache) Del(
	ctx context.Context,
	ns string,
	key string,
) *redis.IntCmd {
	return c.client.Del(ctx, prefix(ns, key))
}

func (c *Cache) TTL(
	ctx context.Context,
	ns string,
	key string,
) *redis.DurationCmd {
	return c.client.TTL(ctx, prefix(ns, key))
}

func prefix(namespace, key string) string {
	return fmt.Sprintf("%s:%s", namespace, key)
}
