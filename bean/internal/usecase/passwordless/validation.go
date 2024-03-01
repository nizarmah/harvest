package passwordless

import (
	"fmt"
	"net/mail"

	"github.com/whatis277/harvest/bean/internal/entity/model"
)

func validateEmail(email string) model.UserInputError {
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
