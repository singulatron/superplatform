---
sidebar_position: 40
tags:
  - user-svc
  - permissions
  - roles
  - authentication
  - authorization
  - services
---

# Dynamic Svc

The dynamic service is designed to help build backendless applications: the goal is to be able to save and query data directly from the frontend. Similarly to Firebase.

> This page is a high level overview of the `Dynamic Svc`. For more details, please see the [Dynamic Svc API documentation](/docs/superplatform/query).

Aimed at prototyping or where building a service to store the data feels like an overkill. It doesn't aim to be a definitive and exclusively used datastore by any means.

> Currently the Dynamic Svc is being used with internal apps but the goal is to have a permission model that works for public apps (where adversarial users might be present). If you find a logical inconsitency that hinders building public apps, please report it.

## How It Works

### Data Model

Multiple tenants (users, services) write to the same table/s. Rows are then owned by whoever created them and access is dictated by the permissions, see Permission Model below.

### Permission Model

The Dynamic Svc has a permission model with the following goals:

- Be simple & easy to understand
- Be as versatile as possible while being simple

To understand the permission model, lets disect an example entry:

```json
{
  "authors": ["usr_12345", "org_67890"],
  "data": {},
  "deleters": ["usr_12345"],
  "id": "pet_67890",
  "readers": ["org_67890"],
  "table": "pet",
  "writers": ["org_67890"]
}
```

### Readers

Readers are user ids, organization ids or role ids that can read the entry.

You can specify other users' IDs or IDs of organizations you are not part of. This can sometimes cause "spam" in multitenant applications where adversarial entities can be present on the same platform. To fix this issue, see the `authors` field.

### Authors

The `authors` field simply marks which user or organization created the entry. This field is used to avoid "spam".

> In certain platforms spam is because anyone can "offer" a record to be read by an other user or organizations they are not part of. Sometimes this behaviour is undesired: imagine a chat application where strangers spam messages just by knowing the company ID. The Authors field fixes this.

### Writers

Writers are user ids, organization ids or role ids that can edit entry.

### Deleters

Deleters are user ids, organization ids or role ids that can delete the entry.

## Conventions

### Table Name and Record ID

The ID of a record must be prefixed by the table name. For this reason, use singular table names if possible.

```
{
  "table": "pet",
  "id": "pet_67890"
}
```
