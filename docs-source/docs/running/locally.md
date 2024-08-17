---
sidebar_position: 1
tags:
  - run
  - deploy
---

# Running Locally

## Docker Compose

The easiest way to run this is to clone the repo, step into the repo root and run:

```sh
docker compose up
```

or

```sh
docker compose up -d
```

to run in the background. The `docker-compose-yaml` in the root folder is designed to build and run the current code. For a more production ready Docker Compose file see the [Docker Compose page](./docker-compose/).

## Natively (Go & Angular)

If you have both Go and Angular installed on your computer, the easiest way to dip your feet into Singulatron is to run things locally.

## Backend

```bash
cd localtron;
go run main.go
```

## Frontend

```bash
cd desktop/workspaces/angular-app/;
npm run start
```

## Administration

### Local files

By default Singulatron uses the folder `~/.singulatron` on your machine for config, file downloads and for the local database.

#### Config file

```bash
cat ~/.singulatron/config.yaml
```

#### Download.json

This file contains all the local downloads on a node. Losing is file is not a big deal as downloaded files are detected even if this file or the entry in this file is missing.

```bash
~/.singulatron/downloads.json
```

#### Data files

By default Singulatron uses local gzipped json files to store database entries. Data access across Singulatron is interface based so the this implementation can be easily swapped out for PostgreSQL and other database backends.

The files are located at

```bash
ls ~/.singulatron/data
```

If you want to view the contents of a file:

```bash
cat ~/.singulatron/data/users.zip | gzip -dc

# or if you jave jq installed
cat ~/.singulatron/data/users.zip | gzip -dc | jq
```
