package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Options struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Cache struct {
	Client *redis.Client
}

func New(opts *Options) (*Cache, error) {
	addr := fmt.Sprintf("%s:%s", opts.Host, opts.Port)

	client := redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Username: opts.Username,
			Password: opts.Password,
		},
	)
	if client == nil {
		return nil, errors.New("error creating redis client")
	}

	err := client.Ping(context.Background()).Err()
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
