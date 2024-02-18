package redistest

import (
	"os"
	"testing"

	"github.com/whatis277/harvest/bean/internal/driver/redis"
)

func CacheTest(t *testing.T) *redis.Cache {
	t.Helper()

	if os.Getenv("INTEGRATION") == "" {
		t.Skip("skipping integration test, set env var INTEGRATION=1")
	}

	cache, err := redis.New(&redis.Config{
		Host:     "redis",
		Port:     "6379",
		Username: "default",
		Password: "",
	})

	if err != nil {
		t.Fatalf("cache error: %s", err)
	}

	t.Cleanup(func() {
		cache.Close()
	})

	return cache
}
