---
title: Docker Compose
nav_order: 1
parent: Server Setup
---

# Docker Compose

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

Please keep in mind that this will store data locally on your machine in as gzipped JSON because Singulatron defaults to local file storage.

See the [Environment Variables](./environment_variables.html)

## Using Your Server

Unless you configured otherwise, you can log in with the following default credentials:

```sh
username: singulatron
password: changeme
```
