package env

import (
	"os"
	"testing"
)

func TestLookup(t *testing.T) {
	t.Run("returns value", func(t *testing.T) {
		os.Setenv("key", "value")

		if v, _ := lookup("key"); v != "value" {
			t.Errorf("expected: %s, got: %s", "value", v)
		}
	})

	t.Run("returns error", func(t *testing.T) {
		os.Clearenv()

		if _, err := lookup("var"); err == nil {
			t.Errorf("expected: %s, got: %s", "error", "nil")
		}
	})
}

func TestLookupBool(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		os.Setenv("key", "true")

		if v, _ := lookupBool("key"); v != true {
			t.Errorf("expected: %t, got: %t", true, v)
		}
	})

	t.Run("returns false", func(t *testing.T) {
		os.Setenv("key", "false")

		if v, _ := lookupBool("key"); v != false {
			t.Errorf("expected: %t, got: %t", false, v)
		}
	})
}
