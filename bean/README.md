# bean
It's good for you.

## Description

A subscription tracker for the rest of us.

## Links

### Development

* Bean: https://whatisbean.local
* SMTP: http://localhost:8025

### Production

TODO

### Documentation

* [Migrations](./internal/driver/migration/README.md)
* [Seeds](./internal/driver/seed/README.md)

There are other documentations under `<dir>/README.md` where relevant.

## Architecture

### Development

* Postgres: Database
* Redis: Cache
* MailHog: SMTP server

### Production

* Postgres: Database
* Redis: Cache
* Postmark: SMTP server

## Guides

### Setup

1. Create a `.env` file in `bean/config`

    The defaults from `.env.example` will be used.
    So, it can be empty for now.

    ```bash
    > touch bean/config/.env
    ```

2. Add `whatisbean.local` to `/etc/hosts`

    ```bash
    # whatisbean.com local
    127.0.0.1 whatisbean.local
    ```

### Run

```bash
> docker compose up caddy bean
```

### Tests

```bash
> docker compose exec -e INTEGRATION=1 bean go test ./... -v
```

## Testing

### General flow

1. Visit the landing page at https://whatisbean.local/

2. Click on "Get started" to sign up or login

3. Enter your email address to sign up or login
    * Use `277@hey.com` for existing data

4. Visit mailhog at http://localhost:8025/

5. Open the email sent and click on the auth URL

6. Ensure you're redirected to https://whatisbean.local/home

7. Create a new card

8. Add a new subscription to the card

9. Delete the subscribe you just created

10. Delete the card you just created

11. Click on "Logout" to logout
