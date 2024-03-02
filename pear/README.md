# pear
It's a low-hanging fruit.

## Description

An open-source guide for your benefit.

## Links

### Production

* Pear: https://whatispear.com

### Development

* Pear: http://whatipear.local

## Guide

### Setup

1. Add `whatispear.local` to `/etc/hosts`

    ```bash
    # whatispear.com local
    127.0.0.1 whatispear.local
    ```

### Run

```bash
> docker compose up trellis pear
```

### Deploy

1. Upload [`./html`](https://github.com/whatis277/harvest/tree/main/pear/html) folder to [Netlify](https://app.netlify.com/sites/harvest-pear/deploys).

2. Test the deployment

3. Publish the deployment
