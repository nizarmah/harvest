package env

import (
	"os"
	"testing"
)

func TestLookup(t *testing.T) {
	t.Run("returns var value", func(t *testing.T) {
		os.Setenv("var", "found")

		if v, _ := lookup("var"); v != "found" {
			t.Errorf("expected: %s, got: %s", "found", v)
		}
	})

	t.Run("returns error if var not found", func(t *testing.T) {
		os.Clearenv()

		if _, err := lookup("var"); err == nil {
			t.Errorf("expected: %s, got: %s", "error", "nil")
		}
	})
}
