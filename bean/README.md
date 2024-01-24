# bean
It's good for you.

## Description

A subscription tracker for the rest of us.

## Run

```bash
> docker compose up bean
```

## Test

```bash
> docker compose exec -e INTEGRATION=1 bean go test ./... -v
```

## Migrate

```bash
> docker compose exec bean tern migrate --destination -+1
```

## Seed

```bash
> docker compose exec bean tern migrate --config ./bean/config/.seed.tern.conf --migrations ./bean/internal/driver/seed --destination -+1
```
