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
    # The `BACKEND_ADDRESS` must be accessible from the browser. It is not an internal address that should be accessible by the frontend. It's the address the browser will make API requests to.
    - BACKEND_ADDRESS=http://127.0.0.1:58231

singulatron-backend:
  image: crufter/singulatron-backend
  ports:
    - "58231:58231"
  volumes:
    # we mount the docker socket so the backend can start containers
    - /var/run/docker.sock:/var/run/docker.sock
    # without this no data will be persisted
    - singulatron-data:/root/.singulatron
  environment:
    # this folder will be mounted by the LLM containers to access the models
    - SINGULATRON_HOST_FOLDER=/var/lib/docker/volumes/singulatron-data/_data
    # address of the host so we can access the containers running the LLMs from the backend container
    - SINGULATRON_LLM_HOST=172.17.0.1
    #
    # For NVIDIA GPU acceleration uncomment the following flags
    #
    # - SINGULATRON_GPU_ENABLED=true
    # - SINGULATRON_IMAGE_OVERRIDE=crufter/llama-cpp-python-cuda
```

There are a few things that can go wrong here:

- You might be on a network that's different from the docker default so `172.17.0.1` might not be the correct one. Try `172.18.0.1` etc. or use `docker network inspect` to find out the IP range of the network you use.
- The `BACKEND_ADDRESS` in a publicly accessible setup should be something like `https://singulatron-api.yourdomain.com`. The point is that it must be accessible from the outside.

## Using Your Server

Unless you configured otherwise, you can log in with the following default credentials:

```sh
username: singulatron
password: changeme
```
