package bcrypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type hasher struct{}

func New() *hasher {
	return &hasher{}
}

func (h *hasher) Hash(input string) (string, error) {
	pass := []byte(input)
	if len(pass) > 72 {
		return "", errors.New("input is too long")
	}

	bytes, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing input")
	}

	return string(bytes), nil
}

func (h *hasher) Compare(input, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	if err != nil {
		return errors.New("input does not match hash")
	}

	return nil
}
