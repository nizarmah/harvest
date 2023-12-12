package env

import (
	"fmt"
	"os"
)

func newDB() (*DB, error) {
	name, err := lookup("DB_NAME")
	if err != nil {
		return nil, err
	}

	host, err := lookup("DB_HOST")
	if err != nil {
		return nil, err
	}

	username, err := lookup("DB_USERNAME")
	if err != nil {
		return nil, err
	}

	password, err := lookup("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	return &DB{
		Name:     name,
		Host:     host,
		Username: username,
		Password: password,
	}, nil
}

func lookup(s string) (string, error) {
	v, ok := os.LookupEnv(s)
	if !ok {
		return "", fmt.Errorf("env variable '%s' not found", s)
	}

	return v, nil
}
