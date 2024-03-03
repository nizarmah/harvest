# bean
It's good for you.

## Description

A subscription tracker for the rest of us.

## Links

### Development

* Bean: http://whatisbean.local
* SMTP: http://localhost:8025

### Production

* Bean: https://whatisbean.com
* SMTP: https://app.postmarkapp.com

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
> docker compose up trellis bean
```

### Tests

```bash
> docker compose exec -e INTEGRATION=1 bean go test ./... -v
```

### Deploy

#### Database

On the first run, you need to create the user, database, and schema.

```sql
CREATE USER bean WITH PASSWORD 'secret-password';
CREATE DATABASE bean;
GRANT CONNECT ON DATABASE bean TO bean;

\c bean
CREATE SCHEMA bean AUTHORIZATION bean;
GRANT ALL ON SCHEMA bean TO bean;
```

#### Server

```bash
> basket/compose pull bean
> basket/compose up -d --no-deps --force-recreate bean
```

#### Migrations

```bash
> basket/compose exec bean tern migrate
```

## Testing

### General flow

1. Visit the landing page at http://whatisbean.local/

2. Click on "Login" to login

3. Enter your email address to login
    * Use `277@hey.com` for existing data

4. Visit mailhog at http://localhost:8025/

5. Open the email sent and click on the auth URL

6. Ensure you're redirected to http://whatisbean.local/home

7. Create a new card

8. Add a new subscription to the card

9. Delete the subscribe you just created

10. Delete the card you just created

11. Click on "Logout" to logout
