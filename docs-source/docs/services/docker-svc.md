---
sidebar_position: 20
tags:
  - docker-svc
  - docker
  - containers
  - service
---

# Docker Svc

The docker service maintains AI model containers on a node. It expects the docker socket to be mounted.

> This page is a high level overview of the `Docker Svc`. For more details, please see the [Docker Svc API documentation](/docs/superplatform/launch-container).

## Responsibilities

- Start and stop containers when needed - ensuring the running containers match what is expected.

## Dependencies

- [Download Svc](/docs/services/download-svc) to get the local file path of the model from the asset URL

## Current Limitations

- Service expects to run on the same node as the Download Svc. This is an issue in distributed setups.