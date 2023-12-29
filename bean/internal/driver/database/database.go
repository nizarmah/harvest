package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func New(builder *DSNBuilder) (*DB, error) {
	pool, err := sql.Open("mysql", dsn(builder))
	if err != nil {
		return nil, err
	}

	err = pool.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		Pool: pool,
	}, nil
}

func (ds *DB) Close() error {
	err := ds.Pool.Close()
	if err != nil {
		return err
	}

	return nil
}
