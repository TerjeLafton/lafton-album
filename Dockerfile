FROM golang:latest as builder

LABEL maintainer="Terje Lafton <terje@lafton.io>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cmd/main .

EXPOSE 8080

CMD ["./main"]