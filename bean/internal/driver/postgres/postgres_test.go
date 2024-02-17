package postgres_test

import (
	"context"
	"testing"

	"github.com/whatis277/harvest/bean/internal/driver/postgres/postgrestest"
)

func TestDB(t *testing.T) {
	db := postgrestest.DBTest(t)

	err := db.Pool.Ping(context.TODO())
	if err != nil {
		t.Fatalf("db error: %s", err)
	}
}
