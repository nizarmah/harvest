package datasource

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func New(builder *DsnBuilder) (*DataSource, error) {
	db, err := sql.Open("mysql", dsn(builder))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DataSource{
		db: db,
	}, nil
}

func (ds *DataSource) Close() error {
	err := ds.db.Close()
	if err != nil {
		return err
	}

	return nil
}
