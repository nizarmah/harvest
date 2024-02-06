package passwordless

import (
	"strings"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tests := []string{
			"277@hey.com",
			"bean@whatis277.com",
			"nizar@whatis277.com",
			"nizarmah@hotmail.com",
			"nizar.mah99@gmail.com",
		}

		for _, test := range tests {
			if err := validateEmail(test); err != nil {
				t.Errorf("expected nil, got: %s", err)
			}
		}
	})

	t.Run("invalid", func(t *testing.T) {
		tests := []string{
			"a",
			"aa",
			strings.Repeat("a", 256),
			"bean",
			"bean@",
			"@bean",
		}

		for _, test := range tests {
			if err := validateEmail(test); err == nil {
				t.Errorf("expected error, got nil: %s", test)
			}
		}
	})
}
