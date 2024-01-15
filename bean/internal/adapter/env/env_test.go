package env

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("DB_NAME", "db_name")
	os.Setenv("DB_HOST", "db_host")
	os.Setenv("DB_USERNAME", "db_username")
	os.Setenv("DB_PASSWORD", "db_password")

	env, err := New()

	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if env.DB.Name != "db_name" {
		t.Errorf("expected: %s, got: %s", "db_name", env.DB.Name)
	}

	if env.DB.Host != "db_host" {
		t.Errorf("expected: %s, got: %s", "db_host", env.DB.Host)
	}

	if env.DB.Username != "db_username" {
		t.Errorf("expected: %s, got: %s", "db_username", env.DB.Username)
	}

	if env.DB.Password != "db_password" {
		t.Errorf("expected: %s, got: %s", "db_password", env.DB.Password)
	}
}
