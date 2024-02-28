package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, builder *DSNBuilder) (*DB, error) {
	pool, err := pgxpool.New(ctx, dsn(builder))
	if err != nil {
		return nil, errors.New("error creating pool")
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, errors.New("error pinging pool")
	}

	return &DB{
		Pool: pool,
	}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}
