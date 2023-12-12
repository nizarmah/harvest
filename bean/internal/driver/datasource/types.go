package datasource

import (
	"database/sql"
)

type DsnBuilder struct {
	Host        string
	Name        string
	Username    string
	Password    string
	Tls         bool
	Interpolate bool
}

type DataSource struct {
	db *sql.DB
}
