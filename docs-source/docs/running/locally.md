---
sidebar_position: 1
tags:
  - run
  - deploy
  - local
---

# Running Locally

The easiest way to run Singulatron is to use [Docker Compose](https://docs.docker.com/compose/install/).

## Docker Compose

The easiest way to run this is to clone [the repo](https://github.com/singulatron/singulatron), step into the repo root and run:

```sh
git clone git@github.com:singulatron/singulatron.git
cd singulatron
docker compose up
# or use the -d flag to run it in the background
# docker compose up -d
```

The `docker-compose-yaml` in the root folder is designed to build and run the current code. For a more production ready Docker Compose file see the [Docker Compose page](./docker-compose/).

### Once it's running

After the containers successfully start, you can go to [http://127.0.0.1:3901](http://127.0.0.1:3901) and log in with the [Default Credentials](/docs/running/using#default-credentials).

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

### Once it's running

After the both the backend and frontend starts, you can go to [http://127.0.0.1:4200](http://127.0.0.1:4200) and log in with the [Default Credentials](/docs/running/using#default-credentials).

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
