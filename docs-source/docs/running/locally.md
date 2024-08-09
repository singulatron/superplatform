---
sidebar_position: 1
tags:
  - run
  - deploy
---

# Running Locally

If you have both Go and Angular installed on your computer, the easiest way to dip your feet into Singulatron is to run things locally.

If you don't have those tools installed, you should probably look into the [Docker Compose example](./docker-compose/).

Without further ado:

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
