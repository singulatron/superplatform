---
sidebar_position: 3
tags:
  - test
---

# Backend Environment Variables

## `SINGULATRON_GPU_PLATFORM`

This envar is used to enabel GPU acceleration.
Supported platforms:

- `cuda`

Do not set this if your card doesn't support the given architecture or things will break.

## `SINGULATRON_HOST_FOLDER`

This envar is needed when Singulatron runs as a container next to containers it starts:

```sh
Host
 |
 |-> Singulatron Container
 |-> Container Launched By Singulatron
```

To be able persist data, a host folder must be mounted by all containers.

In our example (`SINGULATRON_HOST_FOLDER=/var/lib/docker/volumes/singulatron-data/_data`) we basically pass down the full path of the `singulatron-data` volume we created in the docker compose file to the containers created by Singulatron.

So cycle goes like this:

- Singulatron container writes to `/root/.singulatron`, which is mounted on host at `/var/lib/docker/volumes/singulatron-data/_data`
- Assets (which are basically downloaded files) will be passed to containers created by Singulatron by mounting files in `/var/lib/docker/volumes/singulatron-data/_data`.

## `SINGULATRON_LLM_HOST`

When Singulatron is running in a container, it needs to know how to address its siblings (other containers it started):

```sh
Host
 |
 |-> Singulatron Container
 |-> Container Launched By Singulatron
```

The `Singulatron Container` uses the envar `SINGULATRON_LLM_HOST` to address `Container Launched By Singulatron`.

Typically this value should be `172.17.0.1` if you are using the default docker network.

If you are using an other network than default, use `docker network inspect` to find out the IP of your docker bridge for that network.
Usually it's going to be `172.18.0.1`.

This envar is not needed if Singulatron runs directly on the host:

```sh
Host With Singulatron
 |
 |-> Container Launched By Singulatron
```

## `SINGULATRON_DB`

You can use this envar to make Singulatron actually use a database instead of local file storage to store data.

### PostgreSQL

```sh
SINGULATRON_DB=postgres
SINGULATRON_DB_SQL_CONNECTION_STRING="postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
```

Naturally, you should change the details of the connection string to reflect your environment.

## `SINGULARON_LOCAL_STORAGE_PATH`

By default the local file storage will place files into `~/.singulatron/data`, but this flag (and other config options) can override that.
