---
sidebar_position: 20
tags:
  - docker-svc
  - docker
  - containers
  - services
---

# Docker Svc

The docker service maintains containers on a node. It expects the docker socket to be mounted.
For simplicity the Docker Svc it is only concerned with the node it resides on.

## Used By

- [Model Svc](/docs/services/download-svc) to launch AI models.
- [Deploy Svc](/docs/services/download-svc) to launch containers to deploy service instances.

> This page is a high level overview of the `Docker Svc`. For more details, please see the [Docker Svc API documentation](/docs/superplatform/launch-container).

## Responsibilities

- Start and stop containers when needed - ensuring the running containers match what is expected.

## Dependencies

- [Download Svc](/docs/services/download-svc) to get the local file path of the model from the asset URL

## Current Limitations

- Service expects to run on the same node as the Download Svc. This is an issue in distributed setups.
