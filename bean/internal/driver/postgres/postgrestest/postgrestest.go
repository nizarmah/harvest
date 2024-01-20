package postgrestest

import (
	"os"
	"testing"

	"harvest/bean/internal/driver/postgres"
)

func DBTest(t *testing.T) *postgres.DB {
	t.Helper()

	if os.Getenv("INTEGRATION_DB") == "" {
		t.Skip("skipping integration test, set env var INTEGRATION_DB=1")
	}

	db, err := postgres.New(&postgres.DSNBuilder{
		Host:     "postgres",
		Port:     "5432",
		Name:     "bean_test",
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
