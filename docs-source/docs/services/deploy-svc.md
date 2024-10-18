---
sidebar_position: 90
tags:
  - deploy-svc
  - deploy
  - containers
  - services
---

# Deploy Svc

The deploy service is responsible of launching containers on whatever infrastructure the Superplatform is running on (eg. [Docker Svc](/docs/services/docker-svc)) and registering them into the [Registry Svc](/docs/services/docker-svc).

It registers services it launches since services are not expected to self register. This is to support services that are not using the Superplatform SDK to build themselves—in other words, Superplatform is designed to be able to run non-Superplatform services too.

> This page is a high level overview of the `Deploy Svc`. For more details, please see the [Deploy Svc API documentation](/docs/superplatform/save-deployment).

## Dependencies

- [Docker Svc](/docs/services/docker-svc) to start containers of services
- [Registry Svc](/docs/services/registry-svc) to start containers of services
