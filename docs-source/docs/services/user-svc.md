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

The user service is at the heart of Singulatron, managing users, tokens, organizations, permissions and more. Each service and human on the Singulatron network has an account in the `User Svc`.

> This page is a high level overview of the `User Svc`. For API documentation, please see the [API documentation](/docs/singulatron/login).

### Managing Credentials

The most important thing about the User Svc is that service (machine) and user (human) accounts look and function the same.

Every service you write needs to [register](/docs/singulatron/register) at startup, or log in with the credentials it saves and manages if it's already regsitered. Just like a human.

You can do this in a few ways:

- Use the [API](/docs/singulatron/register) directly
- Use a language specific [client](https://github.com/singulatron/singulatron/tree/main/clients) that was generated from the API
- Use a language specific [SDK](https://github.com/singulatron/singulatron/tree/main/localtron/sdk)

#### A Practical Example

A code snippet is worth a thousand words, even if it's in an unfamiliar language, so here is how the Prompt Svc boots up:

```go
func (cs *PromptService) Start() error {
	token, err := sdk.RegisterService("prompt-svc", "Prompt Service", cs.router, cs.credentialStore)
	if err != nil {
		return err
	}
	cs.router = cs.router.SetBearerToken(token)

	return cs.registerPermissions()
}
```

## Roles

### Types of Roles

### Static

Static roles, such as `user-svc:admin` and `user-svc:user` defined by the `User Svc` are primarily used for simple role-based access control: in the Singulatron UI and API you can edit static roles to add or remove endpoints a user can call.

> If you are looking at restricting access to endpoints in other ways, you might be interested in: [Policy Svc](/docs/services/policy-svc).

#### Dynamic

Dynamic roles are generated based on specific user-resource associations, offering more flexible permission management compared to static roles.

Dynamic roles look like `user-svc:org:{org_dBZRCej3fo}:admin`. The dynamic values must be surrounded by `{}` symbols. The above example is how organization roles are represented.

These dynamic roles, like static roles are stored in the JWT tokens so it is advisable to keep them to a minimum. The organization example is an apt one here: think about how many GitHub or Google organizations you are part of. Likely even a few dozen are at the most extreme upper limit.

> JWT tokens (and the dynamic they contain) are sent with each request, so try to be efficient with dynamic roles.

### Conventions

Each role created must by prefixed by the slug of the account that created it. Said account becomes the owner of the role and only that account can edit the role.

## Permissions

### Conventions

Each permission created must by prefixed by the slug of the account that created it. Said account becomes the owner of the permission and only that account can add the permission to a role.

> Once you own a permission (by creating it, and it being prefixed by your account slug), you can add it to any role, not just roles owned by you.

Example; let's say your service is `petstore-svc`. Singulatron prefers fine-grained access control, so you are free to define your own permissions, such as `petstore-svc:read` or `petstore-svc:pet:read`.
