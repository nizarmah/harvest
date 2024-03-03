package postgrestest

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/whatis277/harvest/bean/internal/driver/postgres"
)

func DBTest(t *testing.T) *postgres.DB {
	t.Helper()

	if os.Getenv("INTEGRATION") == "" {
		t.Skip("skipping integration test, set env var INTEGRATION=1")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := postgres.New(ctx, &postgres.DSNBuilder{
		Host:     "postgres",
		Port:     "5432",
		Name:     "bean_dev",
		Username: "postgres",
		Password: "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		t.Fatalf("db error: %s", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db
}
