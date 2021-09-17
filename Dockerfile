# syntax=docker/dockerfile:1

FROM golang:1.16 as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /lafton-album cmd/main.go

FROM ubuntu:latest

WORKDIR /

COPY --from=builder /lafton-album /lafton-album

EXPOSE 8080

ENTRYPOINT ["/lafton-album"]