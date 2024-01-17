package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(builder *DSNBuilder) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), dsn(builder))
	if err != nil {
		return nil, errors.New("error creating pool")
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, errors.New("error pinging pool")
	}

	return &DB{
		Pool: pool,
	}, nil
}

func (ds *DB) Close() {
	ds.Pool.Close()
}
