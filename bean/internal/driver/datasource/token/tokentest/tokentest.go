package tokentest

import (
	"testing"

	"harvest/bean/internal/usecase"

	"harvest/bean/internal/driver/datasource/token"
	"harvest/bean/internal/driver/postgres/postgrestest"
)

func TestTokenDS(t *testing.T) usecase.LoginTokenDataSource {
	t.Helper()

	db := postgrestest.DBTest(t)

	return token.New(db)
}
