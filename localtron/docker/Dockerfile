FROM golang:1.22.4 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app
COPY --from=build /app/main .

EXPOSE 58231

CMD ["./main"]