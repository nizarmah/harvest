# trellis
Branching out.

## Description

A router that connects harvest with the world.

## Guides

### Run

```bash
> docker compose up trellis
```

### Deploy

#### SSL certificate

##### First time deployment

The first time you deploy trellis, you need to deploy certbot first.

```bash
> basket/compose pull certbot
> basket/compose run --rm certbot
```

##### Renewal

```bash
> basket/compose run --rm certbot
> basket/compose exec trellis nginx -s reload
```

#### NGINX

Before you do this, make sure you set up the SSL certificates first.
Otherwise, NGINX will fail to start.

```bash
> basket/compose pull trellis
> basket/compose up -d --no-deps --force-recreate trellis
```
