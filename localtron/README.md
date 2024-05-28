# localtron

```sh
crufter@cruftop:~/mono/localtron$ go run main.go 
2024/04/22 11:53:16 Server started on :58231
```

## Structure

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

## Endpoints

### Download

#### Do

```sh
curl -XPOST -d '{"url": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}' 127.0.0.1:58231/download/do
```

#### List

```sh
curl 127.0.0.1:58231/download/list
```