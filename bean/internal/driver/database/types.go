package database

import (
	"database/sql"
)

type DSNBuilder struct {
	Host        string
	Name        string
	Username    string
	Password    string
	Tls         bool
	Interpolate bool
}

type DB struct {
	Pool *sql.DB
}
