<p align="center">
  <img width="150px" src="https://singulatron.com/assets/logo-lighter.svg" />
  <div align="center">
    <span>
      <h1>Singulatron</h1>
    </span>
    <div>Run and develop self-hosted AI apps.</div>
    <div>
      <a href="https://superplatform.ai">superplatform.ai</a> 
    </div>
  </div>
</p>

[![](https://dcbadge.limes.pink/api/server/https://discord.gg/eRXyzeXEvM)](https://discord.gg/eRXyzeXEvM)
![backend build](https://github.com/singulatron/singulatron/actions/workflows/backend-build-github.yaml/badge.svg)
![frontend build](https://github.com/singulatron/singulatron/actions/workflows/frontend-container-build-github.yaml/badge.svg)

Singulatron is a rapid prototyping environment for the AI age. On the surface it looks like a self-hostable ChatGPT inspired app, but if you delve deeper it's more like a Firebase for the AI age.

It aims to serve the hackers who embark on transforming the world with AI over the next decade.

During his 20+ years of building hundreds of different projects, the author was always drawn to building development platforms ([micro](https://github.com/micro/micro), [1backend](https://github.com/1backend/1backend)), ORMs ([gocassa](https://github.com/gocassa/gocassa)) and various other productivity tools ([ok-borg](https://github.com/ok-borg/borg)).
Singulatron is a bit of all of those.

It's an open-source server daemon and web client that lets you self-host and interact with LLMs, as well as a framework and ecosystem for swiftly creating AI-based applications.

## Primary Use Cases
- [x] Execute AI models on a self-hosted Singulatron instance to ensure privacy.
- [x] Develop backendless applications with access to prompting, datastore, and other features in Singulatron (this is what the authors primarily use Singulatron for).
- [ ] Extend the Singulatron backend with custom endpoints written in any language - in other words deploy new apps on Singulatron (this feature is still in development).

## Roadmap

- [x] AI functionality: prompting, prompt queues, threads, download manager
- [x] Streaming, real time updates
- [x] User management: multi-user support, role-based access control
- [x] Support different database backends (local files, SQL and more is coming) and other distributed primitives
- [x] Publish clients for the daemon in different languages
- [ ] Running, scheduling mini-(or not so mini)-apps built on top of Singulatron
- [ ] Many more

![Run](https://singulatron.com/assets/chat.png?refresh=2)

## Run On Your Servers

See [this document](./docs/start/index.md) to help you get started.

## Run On Your Laptop/PC

Download as a binary for your laptop or PC for Windows or Linux from the website: https://singulatron.com/home  
MacOS support is coming soon.

**Note/Troubleshooting**: currently the focus is on server setups. If the app doesn't want to work on your machine, just make sure Docker is running on your system, as the Docker/VM installation is not entirely reliable on every machine yet.

## Stack

It is an Electron application, with Angular on the frontend and Go on the backend. It becomes a simple web app without electron when hosted over the network.

## License

Singulatron is licensed under AGPL-3.0.

## Status

Fairly early phase but there are quite a few installations chugging along nicely already. Why don't you join them?
