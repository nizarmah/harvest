package redis_test

import (
	"context"
	"testing"

	"github.com/whatis277/harvest/bean/internal/driver/redis/redistest"
)

func TestCache(t *testing.T) {
	cache := redistest.CacheTest(t)

	err := cache.Client.Ping(context.TODO()).Err()
	if err != nil {
		t.Fatalf("cache error: %s", err)
	}

	err = cache.Client.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		t.Fatalf("cache error: %s", err)
	}

	val, err := cache.Client.Get(context.Background(), "key").Result()
	if err != nil {
		t.Fatalf("cache error: %s", err)
	}

	if val != "value" {
		t.Fatalf("expected value to be 'value', got %s", val)
	}
}
