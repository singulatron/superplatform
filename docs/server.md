# Server Setup

## Starting Up

### Docker Compose Example

This snippet will give you a quick idea about how to deploy the the frontend and backend containers so they play nicely together:

```yaml
version: "3.8"

volumes:
  singulatron-data:
    name: singulatron-data
    driver: local

singulatron-frontend:
  image: crufter/singulatron-frontend
  ports:
    - "3901:80"
  environment:
    # The `BACKEND_ADDRESS` must be accessible from the browser.
    # It is not an internal address, it's the address the browser will make API requests to.
    - BACKEND_ADDRESS=http://127.0.0.1:58231

singulatron-backend:
  image: crufter/singulatron-backend
  ports:
    - "58231:58231"
  volumes:
    # We mount the docker socket so the backend can start containers
    - /var/run/docker.sock:/var/run/docker.sock
    # We mount a volume so data will be persisted
    - singulatron-data:/root/.singulatron
  environment:
    # This folder will be mounted by the LLM containers to access the models
    - SINGULATRON_HOST_FOLDER=/var/lib/docker/volumes/singulatron-data/_data
    # Address of the host so we can access the containers running the LLMs from the backend container
    # See "System Specific Settings" on this page for more information.
    - SINGULATRON_LLM_HOST=172.17.0.1
    #
    # GPU Acceleration for NVIDIA GPUs
    # Uncomment for non-NVIDIA GPUs.
    #
    - SINGULATRON_GPU_PLATFORM=cuda
```

### System Specific Settings

There are a few things that can go wrong here:

#### `SINGULATRON_LLM_HOST`

You might be on a network that's different from the docker default so `172.17.0.1` might not be the correct one.

Try `172.18.0.1` etc. or use `docker network inspect` to find out the IP range of the network you use.

The reason why this is needed:

```sh
Host
 |
 |-> Singulatron Container
 |-> Container Launched By Singulatron
```

The `Singulatron Container` uses the envar `SINGULATRON_LLM_HOST` to address `Container Launched By Singulatron`.

This envar is not needed if Singulatron runs directly on the host:

```sh
Host With Singulatron
 |
 |-> Container Launched By Singulatron
```

#### `SINGULATRON_HOST_FOLDER`

Similarly to the `SINGULATRON_LLM_HOST`, this envar is needed when Singulatron runs as a container next to containers it starts:

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



#### `BACKEND_ADDRESS`

In a publicly accessible setup should be something like `https://singulatron-api.yourdomain.com`. The point is that it must be accessible from the outside/browser.

## Using Your Server

Unless you configured otherwise, you can log in with the following default credentials:

```sh
username: singulatron
password: changeme
````
