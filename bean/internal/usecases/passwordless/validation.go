package passwordless

import (
	"fmt"
	"net/mail"
)

func validateEmail(email string) error {
	if len(email) < 3 {
		return fmt.Errorf("email must be greater than 2 chars")
	}

	if len(email) > 255 {
		return fmt.Errorf("email must be less than 255 chars")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("email must be valid")
	}

	return nil
}
