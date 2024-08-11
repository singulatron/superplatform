---
sidebar_position: 1
tags:
  - user-svc
  - permissions
  - roles
  - authentication
  - authorization
---

# User Svc

The user service (abbreviated to `User Svc`) is at the heart of Singulatron.
This page aims to give a high level overview about it. For API call level details see the [API documentation](http://localhost:3000/docs/singulatron/login).

## Writing a service that uses the User Svc

### Managing credentials

The first concept to understand is that service (machine) and user (human) accounts look and function the same.

Every service you write needs to [register](http://localhost:3000/docs/singulatron/register) at startup, or log in with the credentials it saves and manages if it's already regsitered. Just like a human.

All of this should happen through the Singulatron API, or a language specific client that was generated from the API, but we also publish some language specific helper functions in the [SDK folder](https://github.com/singulatron/singulatron/tree/main/localtron/sdk) - they are not strictly needed however.

### Defining your permissions

Let's say your service is `petstore-svc`. Singulatron prefers fine-grained access control, so you are free to define your own permissions, such as `petstore-svc:read` or `petstore-svc:pet:read`. The permission names are up to you, but
