<p align="center">
  <img width="150px" src="https://singulatron.com/assets/logo_circled_grey.svg?v=1" />
  <div align="center">
    <span>
      <h1>Superplatform</h1>
    </span>
    <div>A microservices platform focused on distributed AI management and development.</div>
    <div>
      <a href="https://superplatform.ai">superplatform.ai</a> 
    </div>
  </div>
</p>

[![](https://dcbadge.limes.pink/api/server/https://discord.gg/eRXyzeXEvM)](https://discord.gg/eRXyzeXEvM)
![backend build](https://github.com/singulatron/superplatform/actions/workflows/backend-build-github.yaml/badge.svg)
![frontend build](https://github.com/singulatron/superplatform/actions/workflows/frontend-container-build-github.yaml/badge.svg)

The Superplatform server and ecosystem enables you to self-host AI models, build apps that leverage those models in any language, and utilize a microservices-based communal backend designed to support a diverse range of projects.

## Primary Use Cases

### Run AI Models [x]

Run open-source AI models privately on your own infrastructure, ensuring that your data and operations remain fully under your control.

### Quickly Build Backendless AI Apps [x]

Build backendless application by using Superplatform as a database and AI prompting API. Like Firebase, but with a focus on AI.

### Develop Microservices-Based AI Applications [ ]

Build your own backend services around Superplatform, which itself is built on a microservices architecture. Run these services outside or inside the Superplatform platform.

### Deploy Third-Party AI Apps Easily [ ]

Superplatform is designed to make deploying third-party AI applications straightforward. With its focus on virtualization and containers (primarily Docker) and a microservices, API-first approach (using OpenAPI), Superplatform seamlessly integrates other applications into its ecosystem.

<p align="center">
  <a href="https://singulatron.com/assets/chat.png?refresh=3" target="_blank">
    <img width="200px" src="https://singulatron.com/assets/thumbnail/chat.png?refresh=3" />
  </a>
  <a href="https://singulatron.com/assets/model-explorer.png?refresh=3" target="_blank">
    <img width="200px" src="https://singulatron.com/assets/thumbnail/model-explorer.png?refresh=3" />
  </a>
  <a href="https://singulatron.com/assets/permissions.png?refresh=3" target="_blank">
    <img width="200px" src="https://singulatron.com/assets/thumbnail/permissions.png?refresh=3" />
  </a>
  <a href="https://singulatron.com/assets/api.png?refresh=3" target="_blank">
    <img width="200px" src="https://singulatron.com/assets/thumbnail/api.png?refresh=3" />
  </a>
</p>

## Run On Your Servers

See [this page](https://superplatform.ai/docs/category/running) to help you get started.

## Main Services

https://superplatform.ai/docs/category/singulatron-api

### Prompt Svc

Sync or async, streamed or request/response prompting. Dump tens of thousands of prompts into the system and let the prompt queue process it.

### Download Svc

Start, pause, restart large file downloads to enable you to manage your AI models.

### Chat Svc

Manage threads, messages, assets (generated images etc.).

### User Svc

Role-based access control, JWT, fine tuned permissions and a unique ownership model that lets you build microservices based AI apps.

### Node Svc

Get information about your nodes and GPUs like temperature, VRAM usage etc.

### Dynamic svc

A dynamically typed (schema-free) datastore service that lets you piggyback on Superplatform's database. Lets you build backendless applications on top of Singulatorn. Think Firebase but for AI.

## Run On Your Laptop/PC

We have temporarily discontinued the distribution of the desktop version. Please refer to this page for alternative methods to run the software.

## Stack

It is an Electron application, with Angular on the frontend and Go on the backend. It becomes a simple web app without electron when hosted over the network.

## License

Superplatform is licensed under AGPL-3.0.

## Status

Fairly early phase but there are quite a few installations chugging along nicely already. Why don't you join them?
