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

```
> docker compose exec bean tern migrate --destination -+1
```
