package env

import (
	"fmt"
	"os"
)

func lookup(s string) (string, error) {
	v, ok := os.LookupEnv(s)
	if !ok {
		return "", fmt.Errorf("env variable '%s' not found", s)
	}

	return v, nil
}

func lookupBool(s string) (bool, error) {
	v, err := lookup(s)
	return v == "true", err
}
