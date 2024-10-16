---
sidebar_position: 50
tags:
  - model-svc
  - models
  - ai
  - services
---

# Model Svc

The model service can start, stop AI models across multiple runtimes (eg. Docker) and maintains a database of available models on the platform.

> This page is a high level overview of the `Model Svc`. For more details, please see the [Model Svc API documentation](/docs/superplatform/start-default-model).

## Responsibilities

- Start and stop models
- Maintain a database of models and other related information such as the default model

## Dependencies

- [Docker Svc](/docs/services/docker-svc) to start containerized AI models (eg. Llama, Stabel Diffusion etc.)

## Current Limitations

- Stop model endpoint is missing
