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

## `SINGULATRON_VOLUME_NAME`

This envar is needed when Singulatron runs as a container next to containers it starts:

```sh
Host
 |
 |-> Singulatron Container
 |-> Container Launched By Singulatron
```

For the containers like `llama-cpp` to be able to read the models downloaded by Singulatron we they must both mount the same docker volume.

An example of this can be seen in the root `docker-compose.yaml` file: `SINGULATRON_VOLUME_NAME=singulatron-data`.

So cycle goes like this:

- Singulatron container writes to `/root/.singulatron`, which is mounted to the volume `singulatron-data`
- Assets (which are basically downloaded files) will be passed to containers created by Singulatron by mounting files in `singulatron-data`.

## `SINGULATRON_LLM_HOST`

**This flag is usually not needed as Singulatron gets the IP of the Docker bridge.**

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
