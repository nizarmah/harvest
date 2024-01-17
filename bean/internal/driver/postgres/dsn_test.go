package database

import (
	"testing"
)

func TestDSN(t *testing.T) {
	expected := "postgres://user:password@localhost:5432/database?sslmode=require"

	dsn := dsn(&DSNBuilder{
		Host:     "localhost",
		Port:     "5432",
		Name:     "database",
		Username: "user",
		Password: "password",
		SSLMode:  "require",
	})

	if dsn != expected {
		t.Errorf("expected: %v, got: %v", expected, dsn)
	}
}
