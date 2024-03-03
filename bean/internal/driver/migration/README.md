# migration
Migrations for bean.

On production, they are run manually.

On development, they are run on startup.

## Dependencies

### Config

* [`migration.tern.conf`](../../../config/migration.tern.conf)

### Third parties

* [tern](https://pkg.go.dev/github.com/jackc/tern/v2)

## Guide

### New migration

```bash
docker compose exec bean tern new migration_name
```

### Run migration

```bash
docker compose exec bean tern migrate --destination +1
```

If you want to run a migration up and down, just use `--destination -+1`.
