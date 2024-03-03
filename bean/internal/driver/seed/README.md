# seed
Seeds for development and testing, not production.

## Dependencies

### Config

* [`seed.tern.conf`](../../../config/seed.tern.conf)

### Third parties

* [tern](https://pkg.go.dev/github.com/jackc/tern/v2)

## Guide

### New migration

```bash
docker compose exec bean tern new seed_name --migrations ./bean/internal/driver/seed
```

### Run migration

```bash
docker compose exec bean tern migrate --config ./bean/config/seed.tern.conf --migrations ./bean/internal/driver/seed --destination -+1
```

If you want to run a seed up and down, just use `--destination -+1`.
