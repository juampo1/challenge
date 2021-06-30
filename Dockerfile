FROM golang:alpine

RUN apk add build-base

WORKDIR /app

COPY ./ /app

RUN go mod download

EXPOSE 8080

ENTRYPOINT go run cmd/server.go