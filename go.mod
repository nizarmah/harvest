module harvest

go 1.21.5

require github.com/go-sql-driver/mysql v1.7.1 // direct [bean,]
require github.com/jackc/pgx/v5 v5.5.2 // direct [bean,]

require golang.org/x/crypto v0.18.0 // direct [bean,]

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
