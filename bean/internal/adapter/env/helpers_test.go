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
