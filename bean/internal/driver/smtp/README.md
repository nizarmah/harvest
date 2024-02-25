# smtp
SMTP client for email sending.

## Dependencies

* Mailhog: http://localhost:8025/

## Setup

You need to setup these env variables

```sh
SMTP_HOST=smtp
SMTP_PORT=1025
SMTP_USERNAME=
SMTP_PASSWORD=
```

Make sure `SMTP_USERNAME` and `SMTP_PASSWORD` are empty for development to avoid TLS.
