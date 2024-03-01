package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/whatis277/harvest/bean/internal/driver/redis/redistest"
)

func TestCache(t *testing.T) {
	cache := redistest.CacheTest(t)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := cache.Set(ctx, "test", "key", "value", 0).Err()
	if err != nil {
		t.Fatalf("cache error: %s", err)
	}

	val, err := cache.Get(ctx, "test", "key").Result()
	if err != nil {
		t.Fatalf("cache error: %s", err)
	}

	if val != "value" {
		t.Fatalf("expected value to be 'value', got %s", val)
	}
}
