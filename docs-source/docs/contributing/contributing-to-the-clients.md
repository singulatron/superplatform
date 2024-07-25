---
sidebar_position: 2
tags:
  - contributing
  - clients
---

# Contribute to the Clients

## TypeScript/JavaScript Clients

Making sweeping changes in the clients are hard because of how they depend on each other: `js/types` (`@singulatron/types`) is a dependency of `js/client` (`@singulatron/client`).

### Local workflow

Your local workflow when editing the `@singulatron/types` should be is to issue the `bash link_local.sh` in the `clients/js` folder. The script links up the packages in the correct order for local testing.

### Pull Requests

Once it works locally create a PR and see if `js-client-example.yaml` workflow succeeds. If it does you can proceed with editing the `package.json`s of the libraries and release a new version of them.
