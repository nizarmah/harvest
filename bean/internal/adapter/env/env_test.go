package env

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("BASE_URL", "base_url")

	os.Setenv("DB_NAME", "db_name")
	os.Setenv("DB_HOST", "db_host")
	os.Setenv("DB_PORT", "db_port")
	os.Setenv("DB_USERNAME", "db_username")
	os.Setenv("DB_PASSWORD", "db_password")

	os.Setenv("CACHE_HOST", "cache_host")
	os.Setenv("CACHE_PORT", "cache_port")
	os.Setenv("CACHE_USERNAME", "cache_username")
	os.Setenv("CACHE_PASSWORD", "cache_password")

	os.Setenv("SMTP_HOST", "smtp_host")
	os.Setenv("SMTP_PORT", "smtp_port")
	os.Setenv("SMTP_USERNAME", "smtp_username")
	os.Setenv("SMTP_PASSWORD", "smtp_password")

	env, err := New()

	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if env.BaseURL != "base_url" {
		t.Errorf("expected: %s, got: %s", "base_url", env.BaseURL)
	}

	if env.DB.Name != "db_name" {
		t.Errorf("expected: %s, got: %s", "db_name", env.DB.Name)
	}

	if env.DB.Host != "db_host" {
		t.Errorf("expected: %s, got: %s", "db_host", env.DB.Host)
	}

	if env.DB.Port != "db_port" {
		t.Errorf("expected: %s, got: %s", "db_port", env.DB.Port)
	}

	if env.DB.Username != "db_username" {
		t.Errorf("expected: %s, got: %s", "db_username", env.DB.Username)
	}

	if env.DB.Password != "db_password" {
		t.Errorf("expected: %s, got: %s", "db_password", env.DB.Password)
	}

	if env.Cache.Host != "cache_host" {
		t.Errorf("expected: %s, got: %s", "cache_host", env.Cache.Host)
	}

	if env.Cache.Port != "cache_port" {
		t.Errorf("expected: %s, got: %s", "cache_port", env.Cache.Port)
	}

	if env.Cache.Username != "cache_username" {
		t.Errorf("expected: %s, got: %s", "cache_username", env.Cache.Username)
	}

	if env.Cache.Password != "cache_password" {
		t.Errorf("expected: %s, got: %s", "cache_password", env.Cache.Password)
	}

	if env.SMTP.Host != "smtp_host" {
		t.Errorf("expected: %s, got: %s", "smtp_host", env.SMTP.Host)
	}

	if env.SMTP.Port != "smtp_port" {
		t.Errorf("expected: %s, got: %s", "smtp_port", env.SMTP.Port)
	}

	if env.SMTP.Username != "smtp_username" {
		t.Errorf("expected: %s, got: %s", "smtp_username", env.SMTP.Username)
	}

	if env.SMTP.Password != "smtp_password" {
		t.Errorf("expected: %s, got: %s", "smtp_password", env.SMTP.Password)
	}
}
