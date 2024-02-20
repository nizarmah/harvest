module github.com/whatis277/harvest

go 1.22.0

require github.com/jackc/pgx/v5 v5.5.3 // direct [bean,]

require github.com/redis/go-redis/v9 v9.5.0 // direct [bean,]

require golang.org/x/crypto v0.19.0 // direct [bean,]

require github.com/google/uuid v1.6.0 // direct [bean,]

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
