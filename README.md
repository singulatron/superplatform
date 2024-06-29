<p align="center">
  <img width="150px" src="https://singulatron.com/assets/logo-lighter.png" />
  <div align="center">
    <span>
      <h1>Singulatron</h1>
    </span>
    <div>
      The Self-Hosted AI Superplatform
    </div>
    <div>
      <a href="https://singulatron.com/home">singulatron.com</a>
    </div>
  </div>
<p>
<br />
<br />

Singulatron is the AI Superplatform that runs on your computer(s) and server(s). 
It uses no third party APIs, and you have compelte control over your data and privacy.

## Roadmap

- [x] AI functionality: prompting, threads, prompt queues, download manager
- [x] User management: multi-user support, role-based access control
- [ ] <- [IN PROGRESS] Supporting different database backends (local files, SQL and many more) and other distributed primitives
- [ ] Chat with other users in your organization in Singulatron
- [ ] Running, scheduling mini-(or not so mini)-apps built on top of Singulatron
- [ ] Improving this roadmap : )

![Run](https://singulatron.com/assets/chat.png?refresh=1)
## Run On Your Servers

See [this document](./docs/server.md) to help you get started.

## Run On Your Laptop/PC

Download as a binary for your laptop or PC for Windows or Linux from the website: https://singulatron.com/home  
MacOS support is coming soon.

**Note/Troubleshooting**: currently the focus is on server setups. If the app doesn't want to work on your machine, just make sure Docker is running on your system, as the Docker/VM installation is not entirely reliable on every machine yet.

## Why

We bought quite a few beefy GPUs for our servers but realized we need good software to be able to experiment quickly.
Singulatron aims to be both a desktop app for local usage and also to work as a distributed daemon to drive servers, with a web app frontend client that is the same as the local app.

## Highlights

- [Private](./docs/privacy.md): your chats never leave your computer. Works even without an internet connection
- User management with role based access control: control who can do what in your self hosted installation
- Real-time and fast: utilize your hardware and your time to their full extent
- The prompt queue system lets you input many prompts at once - even across threads - they will be processed sensibly. You can leave threads and return - streaming won't be interrupted
- A download manager makes sure your models are well kept
- Run as a binary (exe, deb etc) locally, or on your servers

## Stack

It is an Electron application, with Angular on the frontend and Go on the backend. It becomes a simple web app without electron when hosted over the network.

## License

Singulatron is dual-licensed under the AGPL-3.0-or-later and a commercial license.

### AGPL-3.0-or-later (Personal Use)

For personal, non-commercial use, this project is licensed under the AGPL-3.0-or-later. See the [LICENSE-PERSONAL-USE](LICENSE-PERSONAL-USE) file for details.

### Commercial License

For commercial use, you must purchase a commercial license. Please refer to the [LICENSE-COMMERCIAL-USE](LICENSE-COMMERCIAL-USE) and [AUTHORS](AUTHORS) file for details.

For more information, visit https://singulatron.com/home.

## Status

Fairly early phase but there are quite a few installations chugging along nicely already. Why don't you join them?
