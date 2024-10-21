---
sidebar_position: 80
tags:
  - registry-service
  - registry
  - microservices
  - addresses
  - authentication
---

# Registry Svc

The registry service is designed to maintain a database of services, service instances and nodes.

Its responsibilities include gathering information about:

- Nodes: each Superplatform server registers itself as a node, which roughly correlates to a physical machine

> This page is a high level overview of the `Registry Svc`. For more details, please see the [Registry Svc API documentation](/docs/superplatform/register-instance).

## Entities

### Definition

A `Definition` or service definition consists of the following things:

- An ID
- A container image to run and some additional information like the internal port exposed etc.
- A set of endpoint definitions (OpenAPI etc.)
- The URL of different clients (JS, Go etc.)

A `Definition` is an abstract concept that can not be called. For a callable entity look at `Instance`s.

### Instance

A `Instance` or a service instance is an actual running, callable instance of a `Definition`.

A `Instance` consists of the following things:

- Address information that can be used to internally address the `Instance`.
- A Deployment ID.

### Deployment

Definitions become instances through the [Deployment entity of the Deploy Service](docs/services/deploy-svc).

### Node

A `Node` is a physical or virtual machine that runs a Singulatron daemon. The daemon can then lauch service instances or other processes such as containers on these machines.

Maintaining a list of nodes is important so the daemon can efficiently distribute workload across the nodes.

## How It Works

The registry is needed when you want to call services not included in the Singulatron daemon. You can think of the daemon as the standard library and services in the registry as third party libraries.

When you want to call a service, you can ask the registry to provide you with a list of instance addresses for a service by service slug. Then you can use any of those instance addresses to make a call.

Things like load balancing should be done on the client side at the moment, the damon does not provide a Proxy Svc yet.
