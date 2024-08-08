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

Singulatron is a rapid prototyping environment for the AI age.

It allows you to self-host and work with LLMs and other AI technologies, offering a comprehensive framework and ecosystem for rapidly building AI-driven applications.

## Primary Use Cases

### Run AI Models [x]

Run open-source AI models privately on your own infrastructure, ensuring that your data and operations remain fully under your control.

### Quickly Build Backendless AI Apps [x]

Build backendless application by using Singulatron as a database and AI prompting API. Like Firebase, but with a focus on AI.

### Develop Microservices-Based AI Applications [ ]

Build your own backend services around Singulatron, which itself is built on a microservices architecture. Run these services outside or inside the Singulatron platform.

### Deploy Third-Party AI Apps Easily [ ]

Singulatron is designed to make deploying third-party AI applications straightforward. With its focus on virtualization and containers (primarily Docker) and a microservices, API-first approach (using OpenAPI), Singulatron seamlessly integrates other applications into its ecosystem.

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

See [this page](https://superplatform.ai/docs/category/start) to help you get started.

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
