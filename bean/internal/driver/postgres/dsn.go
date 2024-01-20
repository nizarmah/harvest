package postgres

import (
	"fmt"
)

type DSNBuilder struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	SSLMode  string
}

func dsn(builder *DSNBuilder) string {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		builder.Username,
		builder.Password,
		builder.Host,
		builder.Port,
		builder.Name,
		builder.SSLMode,
	)

	return dsn
}
