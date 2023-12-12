package env

import (
	"github.com/joho/godotenv"
)

func Load(path string) error {
	return godotenv.Load(path)
}
