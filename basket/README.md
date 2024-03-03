# basket
From farm to table.

## Description

Harvest's primitive deployment tool.

## Dependencies

### Local

* Docker

### Remote

* Ubuntu Jammy 22.04 LTS x64

## Setup

1. Create a new Ubuntu Jammy 22.04 LTS x64 droplet on DigitalOcean.

1. Add your SSH key to the droplet.

1. Let basket do the setup for you.

    ```bash
    > basket/setup <droplet-ip>
    ```

## Guide

### Build

```bash
> basket/build bean latest 1.0.0
```

### Deploy

```bash
> basket/compose pull <service>
> basket/compose up -d --no-deps --force-recreate <service>
```
