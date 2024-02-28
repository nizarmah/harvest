package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/whatis277/harvest/bean/internal/driver/postgres/postgrestest"
)

func TestDB(t *testing.T) {
	db := postgrestest.DBTest(t)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := db.Pool.Ping(ctx)
	if err != nil {
		t.Fatalf("db error: %s", err)
	}
}
