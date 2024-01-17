package database

import (
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	DBTest(t)
}

func DBTest(t *testing.T) *DB {
	t.Helper()

	if os.Getenv("INTEGRATION_DB") == "" {
		t.Skip("skipping integration test, set env var INTEGRATION_DB=1")
	}

	db, err := New(&DSNBuilder{
		Host:     "localhost",
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
