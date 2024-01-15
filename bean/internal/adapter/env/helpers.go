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
