# server

Localtron is the backend engine of Singulatron, written in Go. It is responsible for running and scheduling prompts, downloading models, auth and many other things.


It serves both the desktop app/exe locally (on a laptop or PC) and the web UI over the network in a client-server setting (for example when installed on premise at a company). The UX will be very similar (minus installing runtimes/dependencies).

## How to start

```sh
crufter@cruftop:~/mono/server$ go run main.go
2024/04/22 11:53:16 Server started on :58231
```

## How it works

### Anatomy of a service

```
main.go
    services
        download
           endpoints
                endpointA.go
                endpointB.go
           types
                types.go
           methodA.go
           methodB.go
```
