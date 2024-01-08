package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type hasher struct{}

func New() *hasher {
	return &hasher{}
}

func (h *hasher) Hash(input string) ([]byte, error) {
	pass := []byte(input)
	if len(pass) > 72 {
		return nil, errors.New("input is too long")
	}

	bytes, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error hashing input")
	}

	return bytes, nil
}
